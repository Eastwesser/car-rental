package controller

import (
	"encoding/json"
	"net/http"

	"github.com/eastwesser/car-rental/backend/internal/usecase"
)

type CarHandler struct {
    carService usecase.CarService
}

func NewCarHandler(carService usecase.CarService) *CarHandler {
    return &CarHandler{carService: carService}
}

func (h *CarHandler) GetCars(w http.ResponseWriter, r *http.Request) {
    filter := r.URL.Query().Get("filter")
    cars := h.carService.GetCars(filter)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cars)
}

type OrderHandler struct {
    orderService usecase.OrderService
}

func NewOrderHandler(orderService usecase.OrderService) *OrderHandler {
    return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
    // Check for empty body first
	if r.Body == nil {
		http.Error(w, "Request body is required", http.StatusBadRequest)
		return
	}

	var order struct {
		Car   string `json:"car"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

    decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	
	if err := decoder.Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    // Validate required fields
	if order.Car == "" || order.Name == "" || order.Phone == "" {
		http.Error(w, "All fields (car, name, phone) are required", http.StatusBadRequest)
		return
	}

    if err := h.orderService.CreateOrder(order.Car, order.Name, order.Phone); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Заказ успешно создан! Мы свяжемся с вами в ближайшее время.",
    })
}
