package interestrate

import (
	"context"
	"database/sql"
)

func (rsc InterestRateResource) getDBTx(ctx context.Context) (*sql.Tx, error) {
	return rsc.DB.BeginTx(ctx, nil)
}

func (rsc InterestRateResource) getDB(ctx context.Context) (InterestRate, error) {

	var result InterestRate
	err := rsc.DB.QueryRowContext(ctx, QueryGetInterestRate).Scan(
		&result.ID,
		&result.Rate,
		&result.Compound,
		&result.EffectiveDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		}
		return result, err
	}

	return result, nil
}

func (rsc InterestRateResource) updateDB(ctx context.Context, tx *sql.Tx, req InterestRate) error {

	_, err := tx.ExecContext(
		ctx,
		QueryUpdateInterestRate,
		req.Rate,
		req.Compound,
		req.EffectiveDate,
	)
	if err != nil {
		return err
	}

	return nil
}
