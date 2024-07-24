package account

import (
	"context"
	"database/sql"
)

type (
	AccountDomainItf interface {
		GetAccountByUserID(ctx context.Context, userID int) (AccountEntity, error)
		UpdateAccountBalance(ctx context.Context, tx *sql.Tx, req AccountEntity) error
	}

	AccountDomain struct {
		resource AccountResourceItf
	}

	AccountResourceItf interface {
		getAccountByUserIDDB(ctx context.Context, userID int) (AccountEntity, error)
		updateAccountBalanceDB(ctx context.Context, tx *sql.Tx, req AccountEntity) error
	}

	AccountResource struct {
		DB *sql.DB
	}
)

func InitDomain(rsc AccountResourceItf) AccountDomain {
	return AccountDomain{
		resource: rsc,
	}
}

func (acc AccountDomain) GetAccountByUserID(ctx context.Context, userID int) (AccountEntity, error) {
	return acc.resource.getAccountByUserIDDB(ctx, userID)
}

func (acc AccountDomain) UpdateAccountBalance(ctx context.Context, tx *sql.Tx, req AccountEntity) error {
	return acc.resource.updateAccountBalanceDB(ctx, tx, req)
}
