package account

import (
	"github.com/azizyogo/bank-app/domain/account"
	"github.com/azizyogo/bank-app/domain/interestrate"
	"github.com/azizyogo/bank-app/domain/transaction"
	"github.com/azizyogo/bank-app/domain/user"
	"github.com/azizyogo/bank-app/pkg/config"
)

type AccountUsecase struct {
	cfg config.Config
	acc account.AccountDomainItf
	u   user.UserDomainItf
	tr  transaction.TransactionDomainItf
	ir  interestrate.InterestRateDomainItf
}

func InitAccountUsecase(
	cfg config.Config,
	accountDomain account.AccountDomain,
	userDomain user.UserDomain,
	transactionDomain transaction.TransactionDomain,
	interestRateDomain interestrate.InterestRateDomain,
) *AccountUsecase {
	return &AccountUsecase{
		cfg: cfg,
		acc: accountDomain,
		u:   userDomain,
		tr:  transactionDomain,
		ir:  interestRateDomain,
	}
}
