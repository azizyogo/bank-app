package interestrate

import "time"

type InterestRate struct {
	ID            int64     `json:"id" db:"id"`
	Rate          float32   `json:"rate" db:"rate"`
	Compound      int16     `json:"compound" db:"compound"`
	EffectiveDate time.Time `json:"effective_date" db:"effective_date"`
}
