package interestrate

type (
	// Request
	UpdateInterestRateRequest struct {
		Rate          float32 `json:"rate"`
		Compound      int16   `json:"compound"`
		EffectiveDate string  `json:"effective_date"`
	}
)
