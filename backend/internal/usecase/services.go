package usecase

import (
	"strings"

	"github.com/eastwesser/car-rental/backend/internal/entity"
	"github.com/eastwesser/car-rental/backend/internal/repository"
)

type CarService interface {
    GetCars(filter string) []entity.Car
}

type OrderService interface {
    CreateOrder(car, name, phone string) error
}

type carService struct{}

func NewCarService() CarService {
    return &carService{}
}

func (s *carService) GetCars(filter string) []entity.Car {
    filter = strings.ToLower(filter)
    if filter == "" || filter == "все марки" {
        return entity.Cars
    }
    
    var result []entity.Car
    for _, car := range entity.Cars {
        if strings.Contains(strings.ToLower(car.Title), filter) || strings.ToLower(car.Brand) == filter {
            result = append(result, car)
        }
    }
    return result
}

type orderService struct {
    repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
    return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(car, name, phone string) error {
    return s.repo.CreateOrder(car, name, phone)
}
