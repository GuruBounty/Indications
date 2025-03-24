package helpers

type ErrorResponse struct {
	Error string `json:"error"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type SuscessResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}
