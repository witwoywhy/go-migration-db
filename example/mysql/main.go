package main

import (
	"go-migration-example/infrastructure"

	m "github.com/witwoywhy/go-migration-db/cmd"
)

func init() {
	infrastructure.InitConfig()
}

func main() {
	infrastructure.InitLog()
	infrastructure.InitDb()
	infrastructure.InitMigration()
	m.Run()
}
