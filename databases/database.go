package databases

import (
	"github.com/jinzhu/gorm"
)

type Database interface {
	DbConnection() *gorm.DB
}

func NewDatabase() Database {
	return newSqlite()
}
