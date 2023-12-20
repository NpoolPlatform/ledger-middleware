package coupon

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"

	statementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/statement"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/withdraw/coupon"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entcouponwithdraw "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/couponwithdraw"
	entledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw/coupon"
)

type updateHandler struct {
	*Handler
	couponwithdraw *ent.CouponWithdraw
}

func (h *updateHandler) checkCouponWithdraw(ctx context.Context) error {
	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			CouponWithdraw.
			Query().
			Where(
				entcouponwithdraw.ID(*h.ID),
				entcouponwithdraw.DeletedAt(0),
			).
			Only(ctx)
		if err != nil {
			return err
		}
		h.couponwithdraw = info
		return nil
	})
}

func (h *updateHandler) createOrUpdateLedger(ctx context.Context, tx *ent.Tx) error {
	key := fmt.Sprintf("%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateLedger,
		h.couponwithdraw.AppID,
		h.couponwithdraw.UserID,
		h.couponwithdraw.CoinTypeID,
	)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	info, err := tx.
		Ledger.
		Query().
		Where(
			entledger.AppID(h.couponwithdraw.AppID),
			entledger.UserID(h.couponwithdraw.UserID),
			entledger.CoinTypeID(h.couponwithdraw.CoinTypeID),
			entledger.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	if info == nil {
		info, err = ledgercrud.CreateSet(tx.Ledger.Create(), &ledgercrud.Req{
			AppID:      &h.couponwithdraw.AppID,
			UserID:     &h.couponwithdraw.UserID,
			CoinTypeID: &h.couponwithdraw.CoinTypeID,
			Incoming:   &h.couponwithdraw.Amount,
		}).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	}

	incoming := h.couponwithdraw.Amount
	stm, err := ledgercrud.UpdateSetWithValidate(
		info,
		&ledgercrud.Req{
			Incoming: &incoming,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateCouponWithdraw(ctx context.Context, tx *ent.Tx) error {
	if _, err := crud.UpdateSet(
		tx.CouponWithdraw.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) createStatement(ctx context.Context, tx *ent.Tx) error {
	ioExtra := fmt.Sprintf(
		`{"CouponWithdrawID":"%v","CouponID":"%v"}`,
		h.couponwithdraw.EntID,
		h.couponwithdraw.CouponID.String(),
	)

	sha := sha256.Sum224([]byte(ioExtra))
	key := fmt.Sprintf("%v:%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateLedgerStatement,
		h.couponwithdraw.AppID,
		h.couponwithdraw.UserID,
		h.couponwithdraw.CoinTypeID,
		hex.EncodeToString(sha[:]),
	)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	ioType := types.IOType_Incoming
	ioSubType := types.IOSubType_RandomCouponCash
	if _, err := statementcrud.CreateSet(
		tx.Statement.Create(),
		&statementcrud.Req{
			AppID:      &h.couponwithdraw.AppID,
			UserID:     &h.couponwithdraw.UserID,
			CoinTypeID: &h.couponwithdraw.CoinTypeID,
			Amount:     &h.couponwithdraw.Amount,
			IOType:     &ioType,
			IOSubType:  &ioSubType,
			IOExtra:    &ioExtra,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateCouponWithdraw(ctx context.Context) (*npool.CouponWithdraw, error) {
	handler := &updateHandler{
		Handler: h,
	}
	if err := handler.checkCouponWithdraw(ctx); err != nil {
		return nil, err
	}
	switch {
	case h.State == nil:
		fallthrough //nolint
	case h.State.String() == handler.couponwithdraw.State:
		fallthrough //nolint
	case *h.State != types.WithdrawState_Approved:
		return h.GetCouponWithdraw(ctx)
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateCouponWithdraw(ctx, tx); err != nil {
			return err
		}
		if err := handler.createStatement(ctx, tx); err != nil {
			return err
		}
		if err := handler.createOrUpdateLedger(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCouponWithdraw(ctx)
}
