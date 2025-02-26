package data

import (
	domain "go-migration-db/domain/db"
	"go-migration-db/infrastructure"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/witwoywhy/go-cores/errs"
	"github.com/witwoywhy/go-cores/logger"
)

type service struct{}

func New() domain.Service[domain.Request] {
	return &service{}
}

func (s *service) Execute(request domain.Request, l logger.Logger) (string, errs.Error) {
	m, err := migrate.NewWithDatabaseInstance(
		infrastructure.Data,
		infrastructure.DataDriverEngine,
		infrastructure.DataDriver,
	)
	if err != nil {
		l.Errorf("failed to new instance data pg: %v", err)
		return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
	}

	switch request.Action {
	case domain.Up:
		if err := m.Up(); err != nil {
			l.Errorf("pg failed to migrate data up: %v", err)
			return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
		}
	case domain.Down:
		if err := m.Down(); err != nil {
			l.Errorf("pg failed to migrate data down: %v", err)
			return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
		}
	}

	return "Migration Data Success", nil
}
