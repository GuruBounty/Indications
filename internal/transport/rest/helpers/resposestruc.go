package helpers

type Result struct {
	Result any `json:"result"`
}

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
