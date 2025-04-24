package helpers_test

import (
	"encoding/json"
	"indication/internal/transport/rest/helpers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorResponse(t *testing.T) {
	tests := []struct {
		name    string
		message string
		status  int
		//expected helpers.ErrorResponse
		expectedBody        helpers.Result
		expectedCode        int
		expectedContentType string
	}{
		{
			name:    "returns Internal Server Error",
			message: "Internal Server Error",
			status:  http.StatusInternalServerError,
			expectedBody: helpers.Result{
				Result: helpers.ErrorResponse{
					Error: "Internal Server Error",
				},
			},

			expectedCode:        500,
			expectedContentType: "application/json",
		},
		{
			name:    "returns Error setting meter indication",
			message: "Error setting meter indication",
			status:  http.StatusInternalServerError,
			expectedBody: helpers.Result{
				Result: helpers.ErrorResponse{
					Error: "Error setting meter indication",
				},
			},
			expectedCode:        500,
			expectedContentType: "application/json",
		},
		{
			name:    "returns Invalid login or password",
			message: "Invalid login or password",
			status:  http.StatusUnauthorized,
			expectedBody: helpers.Result{
				Result: helpers.ErrorResponse{
					Error: "Invalid login or password",
				},
			},
			expectedCode:        401,
			expectedContentType: "application/json",
		},
		{
			name:    "returns Bad Request",
			message: "Bad Request",
			status:  http.StatusBadRequest,
			expectedBody: helpers.Result{
				Result: helpers.ErrorResponse{
					Error: "Bad Request",
				},
			},
			expectedCode:        400,
			expectedContentType: "application/json",
		},
		{
			name:    "returns Bad Request Invalid syntax",
			message: "Invalid syntax: 1111",
			status:  http.StatusBadRequest,
			expectedBody: helpers.Result{
				Result: helpers.ErrorResponse{
					Error: "Invalid syntax: 1111",
				},
			},
			expectedCode:        400,
			expectedContentType: "application/json",
		},
		{
			name:    "returns Meter can't be 0",
			message: "meter can't be 0",
			status:  http.StatusBadRequest,
			expectedBody: helpers.Result{
				Result: helpers.ErrorResponse{
					Error: "meter can't be 0",
				},
			},
			expectedCode:        400,
			expectedContentType: "application/json",
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()

		helpers.ReturnResonse(w, tt.message, tt.status)
		contentType := w.Header().Get("Content-Type")
		// check status code
		assert.Equal(t, tt.expectedCode, w.Code, "expected status code %d, got %d", tt.expectedCode, tt.status)
		// check content type
		assert.Equal(t, tt.expectedContentType, contentType, "expected content type %s, got %s", tt.expectedContentType, contentType)
		// check body
		var response helpers.Result

		err := json.NewDecoder(w.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Fataled to decode response body: %v", err)
		}
		gotErrorResponse, ok := response.Result.(map[string]interface{})
		if !ok {
			t.Fatalf("ReturnResonse() Result field is not a map[string]interface{}: %v", response.Result)
		}
		gotError := gotErrorResponse["error"]
		//assert.Equal(t, tt.expectedBody, response, "expected response %v, got %v", tt.expectedBody, response)
		// expectedErrorResponse, ok := tt.expectedBody.Result.(helpers.ErrorResponse)
		// if !ok {
		// 	t.Fatalf("expectedBody.Result is not of type helpers.ErrorResponse: %v", tt.expectedBody.Result)
		// }
		// if response.Result.(map[string]interface{})["error"] != expectedErrorResponse.Error {
		// 	t.Errorf("ReturnResonse() body = %v, want %v", response, tt.expectedBody)
		// }
		if gotError != tt.expectedBody.Result.(helpers.ErrorResponse).Error {
			t.Errorf("ReturnResonse() body = %v, want %v", gotError, tt.expectedBody.Result.(helpers.ErrorResponse).Error)
		}
	}
}
