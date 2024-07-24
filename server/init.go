package server

import (
	"database/sql"
	"fmt"
	"log"

	accountDmn "github.com/azizyogo/bank-app/domain/account"
	interestRateDmn "github.com/azizyogo/bank-app/domain/interestrate"
	transactionDmn "github.com/azizyogo/bank-app/domain/transaction"
	userDmn "github.com/azizyogo/bank-app/domain/user"
	"github.com/azizyogo/bank-app/pkg/config"
	accountUsecase "github.com/azizyogo/bank-app/usecase/account"
	interestrateUsecase "github.com/azizyogo/bank-app/usecase/interestrate"
	"github.com/azizyogo/bank-app/usecase/sandbox"
	userUsecase "github.com/azizyogo/bank-app/usecase/user"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var (
	db  *sql.DB
	cfg config.Config

	// domain
	accountDomain      accountDmn.AccountDomain
	userDomain         userDmn.UserDomain
	interestRateDomain interestRateDmn.InterestRateDomain
	transactionDomain  transactionDmn.TransactionDomain

	// usecase
	AccountUsecase      *accountUsecase.AccountUsecase
	UserUsecase         *userUsecase.UserUsecase
	InterestRateUsecase *interestrateUsecase.InterestRateUsecase
	SandboxUsecase      *sandbox.SandboxUsecase
)

func Init() error {
	cfg = config.LoadConfig()

	var err error
	db, err = initDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return err
	}

	initLayers(db, cfg)

	return nil
}

func Close() {
	if db != nil {
		db.Close()
	}
}

func initLayers(db *sql.DB, cfg config.Config) {

	// init domain
	accountDomain = accountDmn.InitDomain(
		accountDmn.AccountResource{
			DB: db,
		},
	)

	userDomain = userDmn.InitDomain(
		userDmn.UserResource{
			DB: db,
		},
	)

	interestRateDomain = interestRateDmn.InitDomain(
		interestRateDmn.InterestRateResource{
			DB: db,
		},
	)

	transactionDomain = transactionDmn.InitDomain(
		transactionDmn.TransactionResource{
			DB: db,
		},
	)
	// end of init domain

	// init usecase
	AccountUsecase = accountUsecase.InitAccountUsecase(
		cfg,
		accountDomain,
		userDomain,
		transactionDomain,
		interestRateDomain,
	)

	UserUsecase = userUsecase.InitUserUsecase(
		cfg,
		userDomain,
	)

	InterestRateUsecase = interestrateUsecase.InitInterestRateUsecase(
		interestRateDomain,
	)

	SandboxUsecase = sandbox.InitSandboxUsecase(
		cfg,
		userDomain,
	)
	// end of init usecase
}

func initDB(cfg config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
