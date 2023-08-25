package ledger

import "fmt"

func getStatementExtra(statementID string) string {
	return fmt.Sprintf(`{"RollbackStatementID":"%v"}`, statementID)
}
