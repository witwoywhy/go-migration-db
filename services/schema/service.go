package schema

import (
	"net/http"

	domain "github.com/witwoywhy/go-migration-db/domain/db"
	"github.com/witwoywhy/go-migration-db/infrastructure"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/witwoywhy/go-cores/errs"
	"github.com/witwoywhy/go-cores/logger"
)

type service struct{}

func New() domain.Service[domain.Request] {
	return &service{}
}

func (s *service) Execute(request domain.Request, l logger.Logger) (string, errs.Error) {
	m, err := migrate.NewWithDatabaseInstance(
		infrastructure.Schema,
		infrastructure.SchemaDriverEngine,
		infrastructure.SchemaDriver,
	)
	if err != nil {
		l.Errorf("failed to new instance schema pg: %v", err)
		return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
	}

	switch request.Action {
	case domain.Up:
		if err := m.Up(); err != nil {
			l.Errorf("pg failed to migrate schema up: %v", err)
			return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
		}

	case domain.Down:
		if err := m.Down(); err != nil {
			l.Errorf("pg failed to migrate schema down: %v", err)
			return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
		}
	}

	return "Migration Schema Success", nil
}
