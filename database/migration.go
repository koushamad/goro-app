package database

import (
	"github.com/koushamad/goro-app/app/model"
	"github.com/koushamad/goro-db/app/database/sql"
)

func Migration() {
	var migration []interface{}

	migration = append(migration, &model.Hello{})

	sql.Migrate(migration)
}
