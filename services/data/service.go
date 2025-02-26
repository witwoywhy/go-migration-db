package data

import (
	"net/http"

	domain "github.com/witwoywhy/go-migration-db/domain/db"
	"github.com/witwoywhy/go-migration-db/infrastructure"

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
		l.Errorf("failed to new instance data: %v", err)
		return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "")
	}

	switch request.Action {
	case domain.Up:
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			l.Errorf("failed to migrate data up: %v", err)
			return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "")
		}
	case domain.Down:
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			l.Errorf("failed to migrate data down: %v", err)
			return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "")
		}
	}

	return "Migration Data Success", nil
}
