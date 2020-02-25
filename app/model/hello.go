package model

import (
	"github.com/jinzhu/gorm"
)

type Hello struct {
	gorm.Model
	Title string
	Body string
}

func (h *Hello)Create() {
	SQL.Create(h)
}