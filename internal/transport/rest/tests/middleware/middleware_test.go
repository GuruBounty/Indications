package rest

import (
	"bytes"
	"indication/internal/transport/rest"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoggingMiddleware(t *testing.T) {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// a test request
	req, err := http.NewRequest("GET", "/test", nil)

	assert.NoError(t, err)

	//create a responserecord to record the response
	r := httptest.NewRecorder()

	// capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	// call the middleware
	handler := rest.LoggingMiddleware(mockHandler)
	handler.ServeHTTP(r, req)

	//check the logged output
	logOutput := buf.String()
	assert.Contains(t, logOutput, "[GET] -\n")

	//check timestamp is in correct format
	timestamp := logOutput[:25] //assuming the timestamp is always 25 charactes long.
	_, err = time.Parse(time.RFC1123, timestamp)
	assert.NoError(t, err, "Timestamp should be in RFC1123 format")

}
