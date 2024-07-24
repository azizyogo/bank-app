package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/azizyogo/bank-app/server"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	reqBody := LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := reqBody.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := server.UserUsecase.Login(ctx, reqBody.Username, reqBody.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
