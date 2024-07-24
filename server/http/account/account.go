package account

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/azizyogo/bank-app/common/auth"
	constanta "github.com/azizyogo/bank-app/common/const"
	"github.com/azizyogo/bank-app/common/model"
	"github.com/azizyogo/bank-app/server"

	"github.com/gorilla/mux"
)

func HandleGetBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	claims := ctx.Value(constanta.CLAIMS_CONTEXT_KEY).(*auth.Claims)

	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if claims.UserID != userID && claims.Role <= constanta.ROLE_REGULAR {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	balance, err := server.AccountUsecase.GetBalance(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(balance)
}

func HandleGetEncryptionKey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	claims := ctx.Value(constanta.CLAIMS_CONTEXT_KEY).(*auth.Claims)

	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if claims.UserID != userID && claims.Role <= constanta.ROLE_REGULAR {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	key, err := server.AccountUsecase.GetEncryptionKey(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(key)
}

func HandleDeposit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	claims := ctx.Value(constanta.CLAIMS_CONTEXT_KEY).(*auth.Claims)

	reqBody := model.EncryptedRequestResponse{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if claims.UserID != userID && claims.Role <= constanta.ROLE_REGULAR {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	res, err := server.AccountUsecase.Deposit(ctx, reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func HandleWithdraw(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	claims := ctx.Value(constanta.CLAIMS_CONTEXT_KEY).(*auth.Claims)

	reqBody := model.EncryptedRequestResponse{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if claims.UserID != userID && claims.Role <= constanta.ROLE_REGULAR {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	res, err := server.AccountUsecase.Withdraw(ctx, reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func HandleGetCompoundInterest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	claims := ctx.Value(constanta.CLAIMS_CONTEXT_KEY).(*auth.Claims)

	reqBody := model.EncryptedRequestResponse{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if claims.UserID != userID && claims.Role <= constanta.ROLE_REGULAR {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	res, err := server.AccountUsecase.GetCompoundInterest(ctx, reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
