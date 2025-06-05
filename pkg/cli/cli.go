package cli

import (
	"be-tm/pkg/config"
	"be-tm/pkg/migrations"
	"os"
)

func Init(cfg *config.Config) {
	args := os.Args
	if len(args) > 0 {
		for i := 0; i < len(args); i++ {
			switch args[i] {
			case "migrate":
				migrations.AutoMigrate(cfg)
			}
		}
	}
}
