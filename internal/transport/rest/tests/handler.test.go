package rest_test

import (
	"context"
	"indication/internal/domain"
	//"indication/internal/transport/rest"
	//"testing"

	"github.com/stretchr/testify/mock"
)

type MockIndications struct {
	mock.Mock
}

func (m *MockIndications) GetObjectsByNumLS(ctx context.Context, ls int64) ([]domain.LS, error) {
	args := m.Called(ctx, ls)
	return args.Get(0).([]domain.LS), args.Error(1)
}

// func TestHandler_GetObjectsByNumLS(t *testing.T) {
// 	mockService := new(MockIndications)
// 	handler := rest.NewHandler(mockService)
// }
