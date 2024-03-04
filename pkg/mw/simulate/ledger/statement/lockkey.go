package statement

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

func LockKey(appID, userID, coinTypeID uuid.UUID, extra string) string {
	sha := sha256.Sum224([]byte(extra))
	return fmt.Sprintf("%v:%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateSimulateLedgerStatement,
		appID,
		userID,
		coinTypeID,
		hex.EncodeToString(sha[:]),
	)
}
