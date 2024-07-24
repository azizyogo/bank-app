package transaction

import (
	"context"
	"database/sql"
)

type (
	TransactionDomainItf interface {
		InsertNewTransaction(ctx context.Context, tx *sql.Tx, req TransactionEntity) (int, error)
	}

	TransactionDomain struct {
		resource TransactionResourceItf
	}

	TransactionResourceItf interface {
		insertNewTransactionDB(ctx context.Context, tx *sql.Tx, req TransactionEntity) (int, error)
	}

	TransactionResource struct {
		DB *sql.DB
	}
)

func InitDomain(rsc TransactionResourceItf) TransactionDomain {
	return TransactionDomain{
		resource: rsc,
	}
}

func (transaction TransactionDomain) InsertNewTransaction(ctx context.Context, tx *sql.Tx, req TransactionEntity) (int, error) {
	return transaction.resource.insertNewTransactionDB(ctx, tx, req)
}
