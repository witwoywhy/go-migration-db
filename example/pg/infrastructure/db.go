package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/viper"
	"github.com/witwoywhy/go-cores/dbs"
	m "github.com/witwoywhy/go-migration-db/infrastructure"
)

func InitDb() {
	var config dbs.DbConfig

	if err := viper.UnmarshalKey("db", &config); err != nil {
		panic(fmt.Errorf("failed to load config db.first: %v", err))
	}

	db, err := sql.Open("postgres", config.Dsn)
	if err != nil {
		panic(fmt.Errorf("failed to open pg db: %v", err))
	}

	schema, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: m.TableSchema,
	})
	if err != nil {
		panic(fmt.Errorf("failed to get schema driver pg: %v", err))
	}

	data, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: m.TableData,
	})
	if err != nil {
		panic(fmt.Errorf("failed to get data driver pg: %v", err))
	}

	m.SchemaDriver = schema
	m.DataDriver = data
}
