package rest_test

import (
	"context"
	"fmt"
	"indication/internal/transport/rest"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockIndicationsService struct {
	mock.Mock
}
type MockUserService struct {
	mock.Mock
}

func (m *MockIndicationsService) SetMeterIndicationByGUID(ctx context.Context, uuid string, meter float32, requestID int64) (int64, error) {
	args := m.Called(ctx, uuid, meter, requestID)
	return args.Get(0).(int64), args.Error(1)
}

func TestHandler_SetMeterIndicationByGuid(t *testing.T) {
	tests := []struct {
		name         string
		token        string
		uuid         string
		meter        string
		requestID    string
		expectedCode int
	}{
		{
			name:         "Should return 200 when the requets is valid",
			token:        "123",
			uuid:         "f47ac10b-58cc-11e4-9803-0242ac120002",
			meter:        "123.45",
			requestID:    "123",
			expectedCode: 200,
		},
	}

	for _, tt := range tests {

		mockService := new(MockIndicationsService)
		//mockUserService := new(MockUserService)
		handler := rest.NewHandler(nil, nil)
		mockService.On("SetMeterIndicationByGUID", mock.Anything, tt.uuid, float32(123.45), int64(0)).Return(int64(1), nil)
		req, err := http.NewRequest("GET", fmt.Sprintf("/api/setMeterIndicationByGuid/%s/%s/%s", tt.uuid, tt.meter, tt.requestID), nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		hand := http.HandlerFunc(handler.SetMeterIndicationByGuid)
		hand.ServeHTTP(rr, req)
		assert.Equal(t, tt.expectedCode, rr.Code)
	}

}
func TestIsValidUUID(t *testing.T) {
	tests := []struct {
		name     string
		uuid     string
		expected bool
	}{
		{
			name:     "Valid UUID v1",
			uuid:     "f47ac10b-58cc-11e4-9803-0242ac120002",
			expected: true,
		},
		{
			name:     "Valid UUID v2",
			uuid:     "000003e8-f52a-21ef-8400-325096b39f47",
			expected: true,
		},
		{
			name:     "Valid UUID v3",
			uuid:     "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			expected: true,
		},
		{
			name:     "Valid UUID v4",
			uuid:     "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			expected: true,
		},
		{
			name:     "Valid UUID v5",
			uuid:     "123e4567-e89b-5abc-8123-456789abcdef",
			expected: true,
		},
		{
			name:     "Incorrect Lenght",
			uuid:     "23e4567-e89b-12d3-a456-42665544000",
			expected: false,
		},
		{
			name:     "Incorrect Format",
			uuid:     "550e8400e29b41d4a716446655440000",
			expected: false,
		},
		{
			name:     "Non Hex Decimal Characters",
			uuid:     "12345678-90ab-cdef-ghij-klmnopqrstuv",
			expected: false,
		},
		{
			name:     "Empty String",
			uuid:     "",
			expected: false,
		},
		{
			name:     "Incorrect variant",
			uuid:     "550e8400-e29b-7123-a123-446655440000",
			expected: false,
		},
	}

	for _, tt := range tests {
		result := rest.CheckUUID(tt.uuid)
		assert.Equal(t, tt.expected, result, "Expected %t for a valid UUID %v", tt.expected, tt.uuid)
	}
}
