package migrations

import (
	"be-tm/internal/user"
	"be-tm/pkg/config"
	"be-tm/pkg/database"
	"log"
	"os"
)

func AutoMigrate(cfg *config.Config) {
	database := database.Init(cfg)
	err := database.AutoMigrate(&user.User{}, &user.RegistrationCode{})
	if err != nil {
		log.Fatalf("Failed migration models: %v", err)
	}
	log.Printf("Successful migration models")
	os.Exit(0)
}
