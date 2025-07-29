package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/eastwesser/car-rental/backend/internal/config"
	_ "github.com/lib/pq"
)

type OrderRepository interface {
    CreateOrder(car, name, phone string) error
}

type orderRepository struct {
    db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
    return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(car, name, phone string) error {
    _, err := r.db.Exec(
        "INSERT INTO orders (car, name, phone) VALUES ($1, $2, $3)",
        car, name, phone,
    )
    if err != nil {
        log.Printf("Error creating order: %v", err)
        return errors.New("failed to create order")
    }
    return nil
}

func InitDB(dbConfig config.DBConfig) (*sql.DB, error) {
    connStr := fmt.Sprintf(
        "user=%s dbname=%s password=%s host=%s sslmode=disable",
        dbConfig.User,
        dbConfig.Name,
        dbConfig.Password,
        dbConfig.Host,
    )
    
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    if err = db.Ping(); err != nil {
        return nil, err
    }

    if _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS orders (
            id SERIAL PRIMARY KEY,
            car TEXT NOT NULL,
            name TEXT NOT NULL,
            phone TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `); err != nil {
        return nil, err
    }

    return db, nil
}
