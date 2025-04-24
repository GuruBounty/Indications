package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"indication/internal/transport/rest/helpers"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary Set Meter Indication by UUID
// @Description Sets a meter indication based on the provided UUID.
// @Tags meters
// @Accept json
// @Produce json
// @Param guid path string true "Meter UUID"
// @Param meter path string true "Meter"
// @Param requestId path string true "Request number"
// @Param Authorization header string true "Token"
// @Security BearerAuth
// @Success 200 {object} helpers.SuscessResponse "Successfully set meter indication"
// @Failure 400 {object} helpers.ErrorResponse "Invalid request parameters"
// @Failure 401 {object} helpers.ErrorResponse "Unauthorized, missing or invalid token"
// @Failure 500 {object} helpers.ErrorResponse "Internal server error"
// @Router /api/setMeterIndicationByGuid/{guid}/{meter}/{requestId} [get]
func (h *Handler) SetMeterIndicationByGuid(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		helpers.ReturnResonse(w, "Missing authorization token", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	uuid := vars["guid"]
	if !isValidUUID(uuid) {
		helpers.ReturnResonse(w, fmt.Sprintf("Invalid syntax: %v", uuid), http.StatusBadRequest)
		return
	}

	meter, err := getSomeFloatFromRequest(r)
	if err != nil {
		helpers.ReturnResonse(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.indicationsService.SetMeterIndicationByGUID(context.TODO(), uuid, meter, 0)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if id == 0 {
		helpers.ReturnResonse(w, "Error setting meter indication", http.StatusInternalServerError)
		return
	}

	// response := helpers.SuscessResponse{
	// 	ID:      strconv.FormatInt(1234, 10),
	// 	Message: "true",
	// }

	response := helpers.Result{
		Result: helpers.SuscessResponse{
			ID:      strconv.FormatInt(id, 10),
			Message: "true",
		},
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func getSomeFloatFromRequest(r *http.Request) (float32, error) {
	vars := mux.Vars(r)
	someFloat := vars["meter"]
	meter, err := strconv.ParseFloat(someFloat, 32)
	//meter, err := strconv.ParseFloat(vars["meter"], 32)

	if err != nil {
		msg := fmt.Sprintf("Invalid syntax: %v", someFloat)
		return 0, errors.New(msg)
	}
	if meter == 0 {
		return 0, errors.New("meter can't be 0")
	}
	return float32(meter), nil
}
func isValidUUID(uuid string) bool {
	// Regular expression to match UUID v1-v5
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	result := r.MatchString(uuid)
	return result
}

// for test
func CheckUUID(uuid string) bool {
	return isValidUUID(uuid)
}
