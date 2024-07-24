package interestrate

import (
	"context"

	"github.com/azizyogo/bank-app/domain/interestrate"
)

func (interestRate *InterestRateUsecase) Get(ctx context.Context) (InterestRateResponse, error) {

	interestDetail, err := interestRate.ir.Get(ctx)
	if err != nil {
		return InterestRateResponse{}, err
	}

	res := InterestRateResponse{
		Rate:          interestDetail.Rate,
		Compound:      interestDetail.Compound,
		EffectiveDate: interestDetail.EffectiveDate,
	}

	return res, nil
}

func (interestRate *InterestRateUsecase) Update(ctx context.Context, req UpdateInterestRateRequest) error {

	// start transaction
	tx, err := interestRate.ir.GetDBTx(ctx)
	if err != nil {
		return err
	}

	// rollback if any errors were found
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	err = interestRate.ir.Update(ctx, tx, interestrate.InterestRate{
		Rate:          req.Rate,
		Compound:      req.Compound,
		EffectiveDate: req.EffectiveDate,
	})
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
