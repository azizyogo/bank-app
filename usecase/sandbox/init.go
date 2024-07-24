package sandbox

import (
	"github.com/azizyogo/bank-app/domain/user"
	"github.com/azizyogo/bank-app/pkg/config"
)

type SandboxUsecase struct {
	cfg config.Config
	u   user.UserDomainItf
}

func InitSandboxUsecase(
	cfg config.Config,
	userDomain user.UserDomain,
) *SandboxUsecase {
	return &SandboxUsecase{
		cfg: cfg,
		u:   userDomain,
	}
}
