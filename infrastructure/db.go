package infrastructure

import (
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/witwoywhy/go-cores/dbs"
)

var (
	SchemaDriver database.Driver
	DataDriver   database.Driver
)

var (
	SchemaDriverEngine string = dbs.Postgres
	DataDriverEngine   string = dbs.Postgres
)
