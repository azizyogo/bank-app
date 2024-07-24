package interestrate

import (
	"encoding/json"
	"log"
	"net/http"

	utilTime "github.com/azizyogo/bank-app/common/util/time"
	"github.com/azizyogo/bank-app/server"
	"github.com/azizyogo/bank-app/usecase/interestrate"
)

func HandleGetInterestRate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ir, err := server.InterestRateUsecase.Get(ctx)
	if err != nil {
		log.Printf("Error getting interest rate: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(ir); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func HandleUpdateInterestRate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	reqBody := UpdateInterestRateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// convert wib to utc
	utcTime, err := utilTime.WibToUTC(reqBody.EffectiveDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = server.InterestRateUsecase.Update(ctx, interestrate.UpdateInterestRateRequest{
		Rate:          reqBody.Rate,
		Compound:      reqBody.Compound,
		EffectiveDate: utcTime,
	})
	if err != nil {
		log.Printf("Error updating interest rate: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(reqBody); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
