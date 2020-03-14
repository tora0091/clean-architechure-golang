package databases

import (
	"github.com/jinzhu/gorm"
)

type sqlite struct{}

func newSqlite() Database {
	return &sqlite{}
}

func (*sqlite) DbConnection() *gorm.DB {
	return nil
}
