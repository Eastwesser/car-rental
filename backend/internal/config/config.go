package config

import "os"

type DBConfig struct {
    User     string
    Password string
    Host     string
    Name     string
}

type ServerConfig struct {
    Port string
}

func LoadDBConfig() DBConfig {
    return DBConfig{
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        Host:     os.Getenv("DB_HOST"),
        Name:     os.Getenv("DB_NAME"),
    }
}

func LoadServerConfig() ServerConfig {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    return ServerConfig{Port: port}
}
