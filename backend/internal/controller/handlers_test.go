package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) CreateOrder(car, name, phone string) error {
	args := m.Called(car, name, phone)
	return args.Error(0)
}

func TestOrderHandler_CreateOrder(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		setupMock      func(*MockOrderService)
		expectedStatus int
		expectCall     bool
	}{
		{
			name: "successful order creation",
			requestBody: map[string]string{
				"car":   "Car 1",
				"name":  "John Doe",
				"phone": "1234567890",
			},
			setupMock: func(m *MockOrderService) {
				m.On("CreateOrder", "Car 1", "John Doe", "1234567890").Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectCall:     true,
		},
		{
			name: "missing car field",
			requestBody: map[string]string{
				"name":  "John Doe",
				"phone": "1234567890",
			},
			setupMock:      func(m *MockOrderService) {},
			expectedStatus: http.StatusBadRequest,
			expectCall:     false,
		},
		{
			name: "missing name field",
			requestBody: map[string]string{
				"car":   "Car 1",
				"phone": "1234567890",
			},
			setupMock:      func(m *MockOrderService) {},
			expectedStatus: http.StatusBadRequest,
			expectCall:     false,
		},
		{
			name: "missing phone field",
			requestBody: map[string]string{
				"car":  "Car 1",
				"name": "John Doe",
			},
			setupMock:      func(m *MockOrderService) {},
			expectedStatus: http.StatusBadRequest,
			expectCall:     false,
		},
		{
			name: "invalid JSON",
			requestBody:    "invalid json",
			setupMock:      func(m *MockOrderService) {},
			expectedStatus: http.StatusBadRequest,
			expectCall:     false,
		},
		{
			name: "service error",
			requestBody: map[string]string{
				"car":   "Car 1",
				"name":  "John Doe",
				"phone": "1234567890",
			},
			setupMock: func(m *MockOrderService) {
				m.On("CreateOrder", "Car 1", "John Doe", "1234567890").Return(errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectCall:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockOrderService := new(MockOrderService)
			tt.setupMock(mockOrderService)

			handler := NewOrderHandler(mockOrderService)

			requestBodyBytes, err := json.Marshal(tt.requestBody)
			require.NoError(t, err)

			req, err := http.NewRequest("POST", "/api/orders", bytes.NewBuffer(requestBodyBytes))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler.CreateOrder(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			
			if tt.expectCall {
				mockOrderService.AssertExpectations(t)
			} else {
				mockOrderService.AssertNotCalled(t, "CreateOrder")
			}
		})
	}
}
