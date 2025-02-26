package infrastructure

import m "github.com/witwoywhy/go-migration-db/infrastructure"

func InitMigration() {
	if AppConfig.Env != "" {
		m.Schema = "file://migrations/schema/" + AppConfig.Env
		m.Data = "file://migrations/data/" + AppConfig.Env
	}
}
