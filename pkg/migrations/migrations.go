package migrations

import (
	"be-tm/internal/user"
	"log"
	"os"

	"gorm.io/gorm"
)

func AutoMigrate(database *gorm.DB) {
	err := database.AutoMigrate(&user.User{}, &user.RegistrationCode{})
	if err != nil {
		log.Fatalf("Failed migration models: %v", err)
	}
	log.Printf("Successful migration models")
	os.Exit(0)
}
