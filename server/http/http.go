package http

import (
	constanta "github.com/azizyogo/bank-app/common/const"
	"github.com/azizyogo/bank-app/middleware"
	"github.com/azizyogo/bank-app/server/http/account"
	"github.com/azizyogo/bank-app/server/http/interestrate"
	"github.com/azizyogo/bank-app/server/http/sandbox"
	"github.com/azizyogo/bank-app/server/http/user"

	"github.com/gorilla/mux"
)

// NewServer initializes the HTTP server with routes
func NewServer() *mux.Router {
	router := mux.NewRouter()

	// Account
	router.HandleFunc("/account/login", user.HandleLogin).Methods("POST")
	router.HandleFunc("/account/{userID}/encryption-key", middleware.Authorize(account.HandleGetEncryptionKey, constanta.ROLE_REGULAR)).Methods("GET")
	router.HandleFunc("/account/{userID}/balance", middleware.Authorize(account.HandleGetBalance, constanta.ROLE_REGULAR)).Methods("GET")
	router.HandleFunc("/account/{userID}/deposit", middleware.Authorize(account.HandleDeposit, constanta.ROLE_REGULAR)).Methods("POST")
	router.HandleFunc("/account/{userID}/withdraw", middleware.Authorize(account.HandleWithdraw, constanta.ROLE_REGULAR)).Methods("POST")
	router.HandleFunc("/account/{userID}/compound-interest", middleware.Authorize(account.HandleGetCompoundInterest, constanta.ROLE_REGULAR)).Methods("GET")

	// under development
	// router.HandleFunc("/transaction/history/{id}", interestrate.HandleTransactionHistory).Methods("GET")

	// Interest rate manager
	router.HandleFunc("/interest-rate", middleware.Authorize(interestrate.HandleGetInterestRate, constanta.ROLE_ADMIN)).Methods("GET")
	router.HandleFunc("/interest-rate", middleware.Authorize(interestrate.HandleUpdateInterestRate, constanta.ROLE_ADMIN)).Methods("PUT")

	// Sandbox
	router.HandleFunc("/sandbox/fe/encrypt", middleware.Authorize(sandbox.HandleEncrypt, constanta.ROLE_ALL)).Methods("POST")
	router.HandleFunc("/sandbox/fe/decrypt", middleware.Authorize(sandbox.HandleDecrypt, constanta.ROLE_ALL)).Methods("POST")

	return router
}
