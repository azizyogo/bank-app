package account

// Request
type (
	DepositRequest struct {
		Amount float64 `json:"amount"`
	}

	WithdrawRequest struct {
		Amount float64 `json:"amount"`
	}

	GetCompoundInterestRequest struct {
		UserID        int `json:"user_id"`
		ExpectedYears int `json:"expected_years"`
	}
)

// Response
type (
	GetBalanceResponse struct {
		UserID  int     `json:"user_id"`
		Balance float64 `json:"balance"`
	}

	GetEncryptionKeyResponse struct {
		UserID        int    `json:"user_id"`
		DecryptionKey string `json:"decryption_key"`
		EncryptionKey string `json:"encryption_key"`
	}

	GetCompoundInterestResponse struct {
		UserID            int     `json:"user_id"`
		CurrentBalance    float64 `json:"current_balance"`
		AccumulateBalance float64 `json:"accumulate_balance"`
		Interest          float64 `json:"interest"`
		Rate              float32 `json:"rate"`
		CompoundInterval  int16   `json:"compound_interval"`
		ExpectedYears     int     `json:"expected_years"`
	}

	MessageResponse struct {
		Message string `json:"msg"`
	}
)
