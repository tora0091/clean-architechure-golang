package databases

import (
	"log"

	"github.com/jinzhu/gorm"

	"clean-architechure-golang/config"
)

type sqlite struct{}

func newSqlite() Database {
	return &sqlite{}
}

func (*sqlite) DbConnection() *gorm.DB {
	db, err := gorm.Open(config.Conf.Database.Driver, config.Conf.Database.SourceName)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
