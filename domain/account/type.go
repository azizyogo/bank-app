package account

type AccountEntity struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	AccountNumber int     `json:"account_number"`
	Balance       float64 `json:"balance"`
}
