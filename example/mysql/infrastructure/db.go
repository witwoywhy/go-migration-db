package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/spf13/viper"
	"github.com/witwoywhy/go-cores/dbs"
	m "github.com/witwoywhy/go-migration-db/infrastructure"
)

func InitDb() {
	var config dbs.DbConfig

	if err := viper.UnmarshalKey("db", &config); err != nil {
		panic(fmt.Errorf("failed to load config db.first: %v", err))
	}

	db, err := sql.Open("mysql", config.ToDsn())
	if err != nil {
		panic(fmt.Errorf("failed to open db: %v", err))
	}

	schema, err := mysql.WithInstance(db, &mysql.Config{
		MigrationsTable: m.TableSchema,
	})
	if err != nil {
		panic(fmt.Errorf("failed to get schema driver: %v", err))
	}

	data, err := mysql.WithInstance(db, &mysql.Config{
		MigrationsTable: m.TableData,
	})
	if err != nil {
		panic(fmt.Errorf("failed to get data driver: %v", err))
	}

	m.SchemaDriver = schema
	m.DataDriver = data
}
