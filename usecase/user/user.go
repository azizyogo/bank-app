package user

import (
	"context"
	"errors"

	"github.com/azizyogo/bank-app/common/auth"
)

func (user *UserUsecase) Login(ctx context.Context, username, password string) (LoginResponse, error) {

	res := LoginResponse{}
	resDB, err := user.u.GetUserByUsername(ctx, username)
	if err != nil {
		return LoginResponse{}, err
	}

	if password != resDB.Password {
		return LoginResponse{}, errors.New("wrong password")
	}

	jwtToken, err := auth.GenerateJWT(resDB.ID, user.cfg.JWTSecret, int8(resDB.Role))
	if err != nil {
		return LoginResponse{}, err
	}

	res = LoginResponse{
		AccessToken:  jwtToken,
		RefreshToken: "", // not developed yet
	}

	return res, nil
}
