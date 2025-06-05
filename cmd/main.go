package main

import (
	"be-tm/internal/modules"
	"be-tm/pkg/core"
	"be-tm/pkg/swagger"
)

// @title be-tm api
// @version 1.0
// @description documentation for be-tm api
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	core := core.Init()
	core.CLI()
	modules.Init(core.RouterApiGroup, core.Database)
	swagger.Init(core.Router)
	core.Run()
}
