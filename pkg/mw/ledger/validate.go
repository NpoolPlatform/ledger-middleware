package ledger

import "fmt"

func (h *Handler) validate() error {
	if h.Spendable != nil && h.Locked != nil {
		return fmt.Errorf("spendable & locked is not allowed")
	}
	if h.Spendable == nil && h.Locked == nil {
		return fmt.Errorf("spendable or locked needed")
	}
	if h.Spendable != nil {
		if h.AppID == nil || h.UserID == nil || h.CoinTypeID == nil {
			return fmt.Errorf("invalid appid or userid or cointypeid")
		}
	}
	return nil
}

