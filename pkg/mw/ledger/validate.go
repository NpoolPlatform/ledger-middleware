package ledger

import "fmt"

func (h *Handler) validate() error {
	if h.Spendable != nil && h.Locked != nil {
		return fmt.Errorf("spendable & locked is not allowed")
	}
	if h.Spendable == nil && h.Locked == nil {
		return fmt.Errorf("spendable or locked needed")
	}
	return nil
}
