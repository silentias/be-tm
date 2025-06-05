package config

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	DB  DBConfig `validate:"required"`
	App AppConfig
}

type DBConfig struct {
	Host     string `validate:"required"`
	Port     string `validate:"required,numeric"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Name     string `validate:"required"`
}

type AppConfig struct {
	Host string
	Port string
}

func LoadConfig() *Config {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatalf(".env file is not found")
	}

	cfg := &Config{
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		App: AppConfig{
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
		},
	}

	if cfg.App.Host == "" {
		cfg.App.Host = "0.0.0.0"
	}
	if cfg.App.Port == "" {
		cfg.App.Port = "8080"
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		log.Fatalf("Config validation error: %v", err)
	}

	return cfg
}
