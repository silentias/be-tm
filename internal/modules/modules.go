package modules

import (
	"be-tm/internal/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(router *gin.RouterGroup, database *gorm.DB) {
	user.InitModule(router, database)
	//task.InitModule(router, database)
	//syncCore.InitModule(router, database)
}
