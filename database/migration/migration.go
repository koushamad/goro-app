package migration

import (
	"github.com/koushamad/goro-app/app/model"
	sql2 "github.com/koushamad/goro-db/app/database/sql"
)

func Migration()  {
	var migration []interface{}

	migration = append(migration, &model.Hello{})

	sql2.Migrate(migration)
}