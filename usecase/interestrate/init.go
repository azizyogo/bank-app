package interestrate

import (
	interestRate "github.com/azizyogo/bank-app/domain/interestrate"
)

type InterestRateUsecase struct {
	ir interestRate.InterestRateDomainItf
}

func InitInterestRateUsecase(
	interestRateDomain interestRate.InterestRateDomain,
) *InterestRateUsecase {
	return &InterestRateUsecase{
		ir: interestRateDomain,
	}
}
