package user

import "gorm.io/gorm"

type Repository interface {
	CreateRegistrationCode(registrionCode *RegistrationCode) error
}

type repository struct {
	database *gorm.DB
}

func NewRepository(database *gorm.DB) Repository {
	return &repository{database: database}
}

func (r *repository) CreateRegistrationCode(registrionCode *RegistrationCode) error {
	return r.database.Create(registrionCode).Error
}
