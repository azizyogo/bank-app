package user

import (
	"github.com/azizyogo/bank-app/domain/user"
	"github.com/azizyogo/bank-app/pkg/config"
)

type UserUsecase struct {
	cfg config.Config
	u   user.UserDomainItf
}

func InitUserUsecase(
	cfg config.Config,
	userDomain user.UserDomain,
) *UserUsecase {
	return &UserUsecase{
		cfg: cfg,
		u:   userDomain,
	}
}
