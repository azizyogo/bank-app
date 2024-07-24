package account

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func (rsc AccountResource) getAccountByUserIDDB(ctx context.Context, userID int) (AccountEntity, error) {

	var result AccountEntity
	err := rsc.DB.QueryRowContext(ctx, QueryGetAccountByUserID, userID).Scan(
		&result.ID,
		&result.UserID,
		&result.AccountNumber,
		&result.Balance,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return AccountEntity{}, nil
		}
		return AccountEntity{}, err
	}

	return result, nil
}

func (rsc AccountResource) updateAccountBalanceDB(ctx context.Context, tx *sql.Tx, req AccountEntity) error {

	_, err := tx.ExecContext(
		ctx,
		QueryUpdateAccountBalance,
		req.Balance,
		req.UserID,
	)
	if err != nil {
		return err
	}

	return nil
}
