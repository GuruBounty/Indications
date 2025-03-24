package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"indication/internal/domain"
	"indication/internal/transport/rest/helpers"

	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Indications interface {
	GetObjectsByNumLS(ctx context.Context, ls int64) ([]domain.LS_Object, error)
	//GetObjectsByNumLS(ctx context.Context, ls int64) ([]domain.LS_Object, error)
	SetMeterIndicationByGUID(ctx context.Context, guid string, meter float32, request int64) (int64, error)
}

type User interface {
	GetByCredentials(ctx context.Context, email string, password string) (domain.User, error)
}

type Handler struct {
	indicationsService Indications
	userService        User
}

func NewHandler(indications Indications, userService User) *Handler {
	return &Handler{
		indicationsService: indications,
		userService:        userService,
	}
}
func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(LoggingMiddleware)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	indications := r.PathPrefix("/api").Subrouter()
	{
		//indications.Use(validateToken)

		//indications.HandleFunc("/id", h.getId).Methods("GET")
		indications.HandleFunc("/getObjectsByNumLS/{ls:[0-9]+}", h.getObjectsByNumLS).Methods("GET")
		indications.HandleFunc("/setMeterIndicationByGuid/{guid}/{meter}/{requestId:[0-9]+}", h.SetMeterIndicationByGuid).Methods("GET")
		r.HandleFunc("/api/auth", h.authentication).Methods("POST")
		indications.Use(validateToken)
	}

	r.HandleFunc("/ping", h.PingHandler).Methods("GET")
	return r
}

// @Summary Get ID and Name Information
// @Description Retrieves the ID and name based on query parameters.
// @Tags user
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Param name query string true "User Name"
// @Param age query string false "User Age (optional)"
// @Success 200 {object} helpers.SuscessResponse "Successful response"
// @Failure 400 {object} helpers.ErrorResponse "Bad request due to missing or unexpected parameters"
// @Router /api/id [get]
func (h *Handler) getId(w http.ResponseWriter, r *http.Request) {

	//Define the request parmaters
	requiredParams := map[string]bool{
		"id":   true,
		"name": true,
		"age":  true,
	}

	//Get all query parameters from the request
	queryParams := r.URL.Query()

	//Check for unexpected parametrs
	for param := range queryParams {
		if !requiredParams[param] {
			helpers.ReturnResonse(w, fmt.Sprintf("Unexpected query parameter: '%s'", param), http.StatusBadRequest)
			return
		}
	}
	//Check for missing parameters
	for param := range requiredParams {
		if queryParams.Get(param) == "" {
			helpers.ReturnResonse(w, fmt.Sprintf("Missing parameter '%s'", param), http.StatusBadRequest)
			return
		}
	}

	id := queryParams.Get("id")
	name := queryParams.Get("name")
	//age := queryParams.Get("age")

	// if id == "" {
	// 	http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
	// 	return
	// }

	response := helpers.SuscessResponse{
		ID:      id,
		Message: name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
