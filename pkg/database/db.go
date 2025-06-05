package database

import (
	"be-tm/pkg/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	parametrs := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.Host,
		cfg.DB.Port,
	)

	database, err := gorm.Open(postgres.Open(parametrs), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database successfully")

	return database
}
