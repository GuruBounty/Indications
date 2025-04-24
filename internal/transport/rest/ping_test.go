package rest_test

import (
	"indication/internal/transport/rest"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	handler := rest.NewHandler(nil, nil)
	request, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatalf("ping request failed: %v", err)
	}
	reqTest := httptest.NewRecorder()
	handler.PingHandler(reqTest, request)

	
	body := reqTest.Body.String()
	assert.Equal(t, http.StatusOK, reqTest.Code)
	assert.Contains(t, body, "pong", "response should contain 'pong'")
	assert.Contains(t, body, time.Now().Format(time.RFC1123)[:10], "response should contain current date")

}
