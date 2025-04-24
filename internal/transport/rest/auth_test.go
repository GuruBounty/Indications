package rest

import (
	"context"
	"encoding/json"
	"indication/internal/domain"

	//"indication/internal/transport/rest"
	"indication/internal/transport/rest/helpers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthUserService struct {
	mock.Mock
}

func (m *MockAuthUserService) GetByCredentials(ctx context.Context, user, password string) (domain.User, error) {
	args := m.Called(ctx, user, password)
	return args.Get(0).(domain.User), args.Error(1)
}

func TestHandler_Authentication(t *testing.T) {
	mockUserService := new(MockAuthUserService)
	handler := NewHandler(nil, mockUserService)

	mockUserService.On("GetByCredentials", mock.Anything, "validuser", "validpassword").Return(domain.User{UserID: 1}, nil)

	req, err := http.NewRequest("POST", "/api/auth", nil)
	assert.NoError(t, err)
	req.SetBasicAuth("validuser", "validpassword")
	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(handler.authentication)
	handlerFunc.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var response helpers.Result
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.Result)

	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	mockUserService.AssertExpectations(t)
}

func TestHandler_AuthenticationCredentials(t *testing.T) {
	tests := []struct {
		name         string
		user         string
		pass         string
		expectedCode int
	}{
		{
			name:         "Should return 401 when the user is empty",
			user:         "",
			pass:         "validpassword",
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "Should return 401 when the password is empty",
			user:         "valid",
			pass:         "",
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "Should return 401 when the user and password are empty",
			user:         "",
			pass:         "",
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "Should return 401 when the user and password are invalid",
			user:         "invaliduser",
			pass:         "invalidpassword",
			expectedCode: http.StatusUnauthorized,
		},
	}

	mockUserService := new(MockAuthUserService)
	handler := NewHandler(nil, mockUserService)
	for _, tt := range tests {
		mockUserService.On("GetByCredentials", mock.Anything, tt.user, tt.pass).Return(domain.User{}, nil)
		req, err := http.NewRequest("POST", "/api/auth", nil)
		assert.NoError(t, err)
		req.SetBasicAuth(tt.user, tt.pass)
		rr := httptest.NewRecorder()
		handlerFunc := http.HandlerFunc(handler.authentication)
		handlerFunc.ServeHTTP(rr, req)
		assert.Equal(t, tt.expectedCode, rr.Code)
		var response helpers.Result
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Result)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	}
}
