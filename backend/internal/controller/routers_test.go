// routers_test.go
package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eastwesser/car-rental/backend/internal/entity"
	"github.com/eastwesser/car-rental/backend/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	// Create mock services with proper expectations
	mockCarService := new(usecase.MockCarService)
	mockCarService.On("GetCars", "").Return([]entity.Car{})

	mockOrderService := new(usecase.MockOrderService)
	// Only expect calls when valid requests are made
	mockOrderService.On("CreateOrder", "test-car", "test-name", "test-phone").Return(nil)

	// Create handlers
	carHandler := NewCarHandler(mockCarService)
	orderHandler := NewOrderHandler(mockOrderService)

	router := NewRouter(carHandler, orderHandler)

	tests := []struct {
		name           string
		method         string
		path           string
		requestBody    string
		expectedStatus int
	}{
		{
			name:           "get cars",
			method:         "GET",
			path:           "/api/cars",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "create order - empty body",
			method:         "POST",
			path:           "/api/orders",
			requestBody:    "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "create order - valid request",
			method:         "POST",
			path:           "/api/orders",
			requestBody:    `{"car":"test-car","name":"test-name","phone":"test-phone"}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "not found",
			method:         "GET",
			path:           "/unknown",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body *bytes.Reader
			if tt.requestBody != "" {
				body = bytes.NewReader([]byte(tt.requestBody))
			} else {
				body = bytes.NewReader(nil)
			}

			req, err := http.NewRequest(tt.method, tt.path, body)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
		})
	}

	// Verify mock expectations
	mockCarService.AssertExpectations(t)
	mockOrderService.AssertExpectations(t)
}
