package interestrate

import "time"

type (
	InterestRateResponse struct {
		Rate          float32   `json:"rate"`
		Compound      int16     `json:"compound"`
		EffectiveDate time.Time `json:"effective_date"`
	}

	UpdateInterestRateRequest struct {
		Rate          float32   `json:"rate"`
		Compound      int16     `json:"compound"`
		EffectiveDate time.Time `json:"effective_date"`
	}
)
