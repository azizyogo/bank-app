package transaction

import (
	"context"
	"database/sql"
)

func (rsc TransactionResource) insertNewTransactionDB(ctx context.Context, tx *sql.Tx, req TransactionEntity) (int, error) {

	var id int
	err := tx.QueryRowContext(
		ctx,
		QueryInsertNewTransaction,
		req.UserID,
		req.Amount,
		req.Type,
	).Scan(
		&id,
	)
	if err != nil {
		return -1, nil
	}

	return id, nil
}
