package model

import (
	"github.com/jinzhu/gorm"
	sql2 "github.com/koushamad/goro-db/app/database/sql"
)

var SQL *gorm.DB

func Build() {
	SQL = sql2.Init()
}