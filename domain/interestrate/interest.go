package interestrate

import (
	"context"
	"database/sql"
)

type (
	InterestRateDomainItf interface {
		GetDBTx(ctx context.Context) (*sql.Tx, error)
		Get(ctx context.Context) (InterestRate, error)
		Update(ctx context.Context, tx *sql.Tx, req InterestRate) error
	}

	InterestRateDomain struct {
		resource InterestRateResourceItf
	}

	InterestRateResourceItf interface {
		getDBTx(ctx context.Context) (*sql.Tx, error)
		getDB(ctx context.Context) (InterestRate, error)
		updateDB(ctx context.Context, tx *sql.Tx, req InterestRate) error
	}

	InterestRateResource struct {
		DB *sql.DB
	}
)

func InitDomain(rsc InterestRateResourceItf) InterestRateDomain {
	return InterestRateDomain{
		resource: rsc,
	}
}

func (irDomain InterestRateDomain) GetDBTx(ctx context.Context) (*sql.Tx, error) {
	return irDomain.resource.getDBTx(ctx)
}

func (irDomain InterestRateDomain) Get(ctx context.Context) (InterestRate, error) {
	return irDomain.resource.getDB(ctx)
}

func (irDomain InterestRateDomain) Update(ctx context.Context, tx *sql.Tx, req InterestRate) error {
	return irDomain.resource.updateDB(ctx, tx, req)
}
