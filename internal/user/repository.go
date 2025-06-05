package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateRegistrationCode(registrionCode *RegistrationCode) error
	Exists(model interface{}, query string, args ...interface{}) (bool, error)
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

// Validation exists

func (r *repository) Exists(model interface{}, query string, args ...interface{}) (bool, error) {
	var count int64
	err := r.database.Model(&model).Where(query, args...).Count(&count).Error
	return count > 0, err
}
