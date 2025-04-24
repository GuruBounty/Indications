package helpers

import (
	"encoding/json"
	"net/http"
)

func ReturnResonse(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(
		Result{
			Result: ErrorResponse{
				Error: message,
			},
		})
}
