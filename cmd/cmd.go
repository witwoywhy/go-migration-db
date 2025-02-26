package cmd

import (
	"flag"
	domain "go-migration-db/domain/db"
	"go-migration-db/services/data"
	"go-migration-db/services/schema"

	"github.com/google/uuid"
	"github.com/witwoywhy/go-cores/apps"
	"github.com/witwoywhy/go-cores/errs"
	"github.com/witwoywhy/go-cores/logs"
)

func Run() {
	l := logs.New(map[string]any{
		apps.TraceID: uuid.NewString(),
		apps.SpanID:  uuid.NewString(),
	})

	action := flag.String("a", "up", "for migration 'up' or 'down'")
	migrate := flag.String("m", "schema", "for migration 'schema' or 'data'")

	flag.Parse()

	a := domain.Action(*action)
	if domain.IsNotAction(a) {
		l.Error("wrong action")
		return
	}

	m := domain.MigrateType(*migrate)
	if domain.IsNotMigrateType(m) {
		l.Error("wrong migrate")
		return
	}

	var response string
	var err errs.Error

	switch m {
	case domain.Schema:
		schema := schema.New()
		response, err = schema.Execute(domain.Request{Action: a}, l)
	case domain.Data:
		data := data.New()
		response, err = data.Execute(domain.Request{Action: a}, l)
	}

	if err != nil {
		l.Error(err)
	} else {
		l.Info(response)
	}
}
