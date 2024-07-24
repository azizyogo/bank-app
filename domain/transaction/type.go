package transaction

import "time"

type TransactionEntity struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Type      int8      `json:"type"` // "deposit" == 1 or "withdrawal" == 2
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
