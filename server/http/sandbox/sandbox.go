package sandbox

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/azizyogo/bank-app/common/auth"
	constanta "github.com/azizyogo/bank-app/common/const"
	"github.com/azizyogo/bank-app/common/model"
	"github.com/azizyogo/bank-app/server"
)

func HandleEncrypt(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	claims := ctx.Value(constanta.CLAIMS_CONTEXT_KEY).(*auth.Claims)

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	encrypted, err := server.SandboxUsecase.SandboxEncrypt(claims.UserID, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := model.EncryptedRequestResponse{
		Data: encrypted,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func HandleDecrypt(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := model.EncryptedRequestResponse{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decrypted, err := server.SandboxUsecase.SandboxDecrypt(ctx, req.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(decrypted)
}
