package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Post struct {
	gorm.Model

	Title   string
	Content string
}
