package rest

import (
	"context"
	"encoding/json"
	"indication/internal/domain"
	"indication/internal/transport/rest/helpers"
	"net/http"
)

// @Summary Authenticate User
// @Description Authenticates the user using Basic Authentication and returns a token.
// @Tags auth
// @Accept json
// @Produce json
// @Security BasicAuth
// @Success 200 {object} helpers.TokenResponse "Successful authentication, returns a token"
// @Failure 401 {object} helpers.ErrorResponse "Unauthorized, invalid login or password"
// @Failure 500 {object} helpers.ErrorResponse "Internal server error, couldn't generate token"
// @Router /api/auth [post]
func (h *Handler) authentication(w http.ResponseWriter, r *http.Request) {

	user, password, ok := r.BasicAuth()

	if user == "" || password == "" {
		helpers.ReturnResonse(w, "Invalid login or password", http.StatusUnauthorized)
		return
	}

	if ok && !checkCredentials(h, user, password) {
		//w.Header().Set("WWW-Authenticate", `Basic realm="api"`)
		helpers.ReturnResonse(w, "Invalid login or password", http.StatusUnauthorized)
		return
	}
	token, err := generateToken(password)
	if err != nil {
		helpers.ReturnResonse(w, "Couldn't generate token", http.StatusInternalServerError)
		return
	}
	tokenresponse := helpers.TokenResponse{
		Token: token,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokenresponse)

}

func checkCredentials(h *Handler, user string, pswd string) bool {
	userReturn, err := h.userService.GetByCredentials(context.TODO(), user, pswd)
	if err != nil || (userReturn == domain.User{}) {
		return false
	}
	return true
}
