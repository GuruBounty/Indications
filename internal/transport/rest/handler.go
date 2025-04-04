package rest

import (
	"context"
	"indication/internal/domain"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Indications interface {
	GetObjectsByNumLS(ctx context.Context, ls int64) ([]domain.LS_Object, error)
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

		indications.HandleFunc("/getObjectsByNumLS/{ls:[0-9]+}", h.getObjectsByNumLS).Methods("GET")
		indications.HandleFunc("/setMeterIndicationByGuid/{guid}/{meter}/{requestId:[0-9]+}", h.SetMeterIndicationByGuid).Methods("GET")
		r.HandleFunc("/api/auth", h.authentication).Methods("POST")
		indications.Use(validateToken)
	}

	r.HandleFunc("/ping", h.PingHandler).Methods("GET")
	return r
}
