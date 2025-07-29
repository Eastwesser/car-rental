package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDBConfig(t *testing.T) {
	tests := []struct {
		name     string
		setup    func()
		expected DBConfig
	}{
		{
			name: "load all env variables",
			setup: func() {
				os.Setenv("DB_USER", "testuser")
				os.Setenv("DB_PASSWORD", "testpass")
				os.Setenv("DB_HOST", "localhost")
				os.Setenv("DB_NAME", "testdb")
			},
			expected: DBConfig{
				User:     "testuser",
				Password: "testpass",
				Host:     "localhost",
				Name:     "testdb",
			},
		},
		{
			name:     "missing env variables",
			setup:    func() {},
			expected: DBConfig{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Сохраняем текущие env переменные
			originalUser := os.Getenv("DB_USER")
			originalPassword := os.Getenv("DB_PASSWORD")
			originalHost := os.Getenv("DB_HOST")
			originalName := os.Getenv("DB_NAME")
			defer func() {
				os.Setenv("DB_USER", originalUser)
				os.Setenv("DB_PASSWORD", originalPassword)
				os.Setenv("DB_HOST", originalHost)
				os.Setenv("DB_NAME", originalName)
			}()

			// Настраиваем тестовые env переменные
			tt.setup()

			// Выполняем тестируемую функцию
			got := LoadDBConfig()

			// Проверяем результаты
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestLoadServerConfig(t *testing.T) {
	tests := []struct {
		name     string
		setup    func()
		expected ServerConfig
	}{
		{
			name: "custom port",
			setup: func() {
				os.Setenv("PORT", "9090")
			},
			expected: ServerConfig{
				Port: "9090",
			},
		},
		{
			name:     "default port",
			setup:    func() {},
			expected: ServerConfig{Port: "8080"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Сохраняем текущую env переменную
			originalPort := os.Getenv("PORT")
			defer func() {
				os.Setenv("PORT", originalPort)
			}()

			// Настраиваем тестовые env переменные
			tt.setup()

			// Выполняем тестируемую функцию
			got := LoadServerConfig()

			// Проверяем результаты
			assert.Equal(t, tt.expected, got)
		})
	}
}
