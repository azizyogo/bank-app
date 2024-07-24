package user

import (
	"context"
	"database/sql"
)

type (
	UserDomainItf interface {
		GetDBTx(ctx context.Context) (*sql.Tx, error)

		GetUserByUserID(ctx context.Context, userID int) (UserEntity, error)
		GetUserByUsername(ctx context.Context, username string) (UserEntity, error)

		GetKeyByUserID(ctx context.Context, userID int) (UserKey, error)
		UpdateKeyByUserID(ctx context.Context, publicKey, privateKey string, userID int) error
	}

	UserDomain struct {
		resource UserResourceItf
	}

	UserResourceItf interface {
		getDBTx(ctx context.Context) (*sql.Tx, error)

		getUserByUserIDDB(ctx context.Context, userID int) (UserEntity, error)
		getUserByUsernameDB(ctx context.Context, username string) (UserEntity, error)

		getKeyByUserIDDB(ctx context.Context, userID int) (UserKey, error)
		updateKeyByUserIDDB(ctx context.Context, publicKey, privateKey string, userID int) error
	}

	UserResource struct {
		DB *sql.DB
	}
)

func InitDomain(rsc UserResourceItf) UserDomain {
	return UserDomain{
		resource: rsc,
	}
}

func (u UserDomain) GetDBTx(ctx context.Context) (*sql.Tx, error) {
	return u.resource.getDBTx(ctx)
}

func (u UserDomain) GetUserByUserID(ctx context.Context, userID int) (UserEntity, error) {
	return u.resource.getUserByUserIDDB(ctx, userID)
}

func (u UserDomain) GetUserByUsername(ctx context.Context, username string) (UserEntity, error) {
	return u.resource.getUserByUsernameDB(ctx, username)
}

func (u UserDomain) GetKeyByUserID(ctx context.Context, userID int) (UserKey, error) {
	return u.resource.getKeyByUserIDDB(ctx, userID)
}

func (u UserDomain) UpdateKeyByUserID(ctx context.Context, publicKey, privateKey string, userID int) error {
	return u.resource.updateKeyByUserIDDB(ctx, publicKey, privateKey, userID)
}
