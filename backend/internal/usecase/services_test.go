package usecase

import (
	"testing"

	"github.com/eastwesser/car-rental/backend/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCarService_GetCars(t *testing.T) {
	// Backup and restore original cars
	originalCars := entity.Cars
	defer func() { entity.Cars = originalCars }()

	// Setup test data
	entity.Cars = []entity.Car{
		{ID: "1", Title: "Toyota Camry", Brand: "toyota"},
		{ID: "2", Title: "Honda Accord", Brand: "honda"},
	}

	service := NewCarService()

	tests := []struct {
		name   string
		filter string
		want   []entity.Car
	}{
		{
			name:   "empty filter",
			filter: "",
			want:   []entity.Car{entity.Cars[0], entity.Cars[1]},
		},
		{
			name:   "filter by brand",
			filter: "toyota",
			want:   []entity.Car{entity.Cars[0]},
		},
		{
			name:   "no matches",
			filter: "bmw",
			want:   []entity.Car{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := service.GetCars(tt.filter)
			assert.Equal(t, tt.want, got)
			assert.NotNil(t, got, "Should never return nil slice")
		})
	}

	// Test nil cars case
	t.Run("nil cars list", func(t *testing.T) {
		entity.Cars = nil
		got := service.GetCars("")
		assert.Equal(t, []entity.Car{}, got)
		assert.NotNil(t, got)
	})
}
