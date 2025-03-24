package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"indication/internal/domain"
	"indication/internal/transport/rest/helpers"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary Get Objects by Num LS
// @Description Retrieves objects based on the provided num ls.
// @Tags objects
// @Accept json
// @Produce json
// @Param num_ls path integer true "Number LS"
// @Param Authorization header string true "Token"
// @Security BearerAuth
// @Success 200 {object} []domain.LS_Object "Successful response"
// @Failure 400 {object} helpers.ErrorResponse "Invalid request parameters"
// @Failure 401 {object} helpers.ErrorResponse "Unauthorized, missing or invalid token"
// @Failure 500 {object} helpers.ErrorResponse "Internal server error"
// @Router /api/getObjectsByNumLS/{num_ls} [get]
func (h *Handler) getObjectsByNumLS(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Missing authorization token", http.StatusUnauthorized)
		return
	}
	ls, err := getSomeIntFromRequest(r)

	if err != nil {
		if errors.Is(err, domain.ErrLSNotFound) {

			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("getObjectByNumLS()  error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	lsObject, err := h.indicationsService.GetObjectsByNumLS(context.TODO(), ls)

	if err != nil {

		helpers.ReturnResonse(w, "Database error", http.StatusInternalServerError)
		return
	}
	if len(lsObject) == 0 {
		helpers.ReturnResonse(w, fmt.Sprintf("LS %v not found", ls), http.StatusNotFound)
		// w.WriteHeader(http.StatusNotFound)
		// w.Header().Set("Content-Type", "application/json")
		// fmt.Fprintf(w, `{"error": "Ls %v not found"}`, ls)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lsObject)

	// works with json file
	// resp, err := h.indicationsService.GetObjectsByNumLS(context.TODO(), ls)
	// if err != nil {
	// 	if errors.Is(err, domain.ErrLSNotFound) {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// 	log.Println("getObjectByNumLS()  error:", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// response, err := json.Marshal(resp)
	// if err != nil {
	// 	log.Println("getObjectByNumLS() error:", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Header().Add("Content-Type", "application/json")
	// w.Write(response)
}

func getSomeIntFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	someInt := vars["ls"]
	//ls := r.URL.Query().Get("ls")
	num, err := strconv.ParseInt(someInt, 10, 64)
	if err != nil {
		return 0, err
	}
	if num == 0 {
		return 0, errors.New("ls can't be 0")
	}
	return num, nil
}
