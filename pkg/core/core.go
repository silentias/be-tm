package core

import (
	"be-tm/pkg/config"
	"be-tm/pkg/database"
	"be-tm/pkg/migrations"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Core struct {
	Cfg            *config.Config
	Router         *gin.Engine
	RouterApiGroup *gin.RouterGroup
	Database       *gorm.DB
}

func Init() *Core {
	cfg := config.LoadConfig()
	router := gin.Default()
	app := &Core{
		Cfg:            cfg,
		Router:         router,
		Database:       database.Init(cfg),
		RouterApiGroup: router.Group("api/v1"),
	}
	return app
}

func (c *Core) CLI() {
	args := os.Args
	if len(args) > 0 {
		for i := 0; i < len(args); i++ {
			switch args[i] {
			case "migrate":
				migrations.AutoMigrate(c.Database)
			}
		}
	}
}

func (c *Core) Run() {
	c.Router.Run(c.Cfg.App.Host + ":" + c.Cfg.App.Port)
}
