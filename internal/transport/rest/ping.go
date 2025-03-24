package rest

import (
	"fmt"
	"net/http"
	"time"
)

func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) {
	var msg = time.Now().Format(time.RFC1123)
	w.Write([]byte(fmt.Sprintf("pong, '%s'", msg)))
}
