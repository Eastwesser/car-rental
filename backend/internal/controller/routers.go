package controller

import (
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
    "net/http"
)

func NewRouter(
    carHandler *CarHandler,
    orderHandler *OrderHandler,
) http.Handler {
    r := chi.NewRouter()
    
    r.Use(middleware.Logger)
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:4200"},
        AllowedMethods:   []string{"GET", "POST"},
        AllowedHeaders:   []string{"Content-Type"},
        AllowCredentials: true,
    }))

    r.Handle("/images/*", 
        http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))) // Fixed missing parenthesis

    r.Get("/api/cars", carHandler.GetCars)
    r.Post("/api/orders", orderHandler.CreateOrder)

    return r
}
