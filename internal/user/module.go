package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitModule(router *gin.RouterGroup, database *gorm.DB) {
	repo := NewRepository(database)
	service := NewService(repo)
	controller := NewController(service)
	handler := NewHandler(controller)

	handler.RegisterRoutes(router)
}
