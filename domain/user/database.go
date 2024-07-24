package user

import (
	"context"
	"database/sql"
)

func (rsc UserResource) getDBTx(ctx context.Context) (*sql.Tx, error) {
	return rsc.DB.BeginTx(ctx, nil)
}

func (rsc UserResource) getUserByUserIDDB(ctx context.Context, userID int) (UserEntity, error) {

	var result UserEntity
	err := rsc.DB.QueryRowContext(ctx, QueryGetUserByUserID, userID).Scan(
		&result.ID,
		&result.Username,
		&result.Password,
		&result.Role,
		&result.PublicKey,
		&result.PrivateKey,
		&result.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return UserEntity{}, nil
		}
		return UserEntity{}, err
	}

	return result, nil
}

func (rsc UserResource) getUserByUsernameDB(ctx context.Context, username string) (UserEntity, error) {

	var result UserEntity
	err := rsc.DB.QueryRowContext(ctx, QueryGetUserByUsername, username).Scan(
		&result.ID,
		&result.Username,
		&result.Password,
		&result.Role,
		&result.PublicKey,
		&result.PrivateKey,
		&result.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return UserEntity{}, nil
		}
		return UserEntity{}, err
	}

	return result, nil
}

func (rsc UserResource) getKeyByUserIDDB(ctx context.Context, userID int) (UserKey, error) {

	var result UserKey
	err := rsc.DB.QueryRowContext(ctx, QueryGetUserKeyByUserID, userID).Scan(
		&result.ID,
		&result.PublicKey,
		&result.PrivateKey,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		}
		return result, err
	}

	return result, nil
}

func (rsc UserResource) updateKeyByUserIDDB(ctx context.Context, publicKey, privateKey string, userID int) error {

	_, err := rsc.DB.ExecContext(
		ctx,
		QueryUpdateKeyByUserID,
		publicKey,
		privateKey,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}
