package rest_test

import (
	"indication/internal/transport/rest"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRouter_NoExistentEndpoint(t *testing.T) {
	handler := rest.NewHandler(nil, nil)
	router := handler.InitRouter()

	request, err := http.NewRequest("GET", "/noexistentEndpoint", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, request)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func Test_getObjectsByNumLS_invalidLS(t *testing.T) {
	handler := rest.NewHandler(nil, nil)
	router := handler.InitRouter()
	router.StrictSlash(true)
	request, err := http.NewRequest("GET", "/api/getObjectsByNumLS/123abc", nil)

	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, request)
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d ", rr.Code)
	}
	// assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestInitRouter_UnsupportedMethod(t *testing.T) {
	handler := rest.NewHandler(nil, nil)
	router := handler.InitRouter()
	tests := []struct {
		name       string
		method     string
		url        string
		statusCode int
	}{
		{
			name:       "Unsupported method POST",
			method:     "POST",
			url:        "/api/getObjectsByNumLS/123",
			statusCode: http.StatusNotFound,
		},
		{
			name:       "Unsupported method PUT",
			method:     "PUT",
			url:        "/api/getObjectsByNumLS/123",
			statusCode: http.StatusNotFound,
		},
		{
			name:       "Method is correct but Number of LS contains letters",
			method:     "GET",
			url:        "/api/getObjectsByNumLS/123abc",
			statusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, nil)
			assert.NoError(t, err)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			//assert.Equal(t, tt.statusCode, rr.Code)
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d; got %d", tt.statusCode, rr.Code)
			}
		})
	}
}
