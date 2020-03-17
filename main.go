package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"clean-architechure-golang/config"
	"clean-architechure-golang/databases"
	"clean-architechure-golang/router"
)

func main() {
	config.ReadConf()

	dbconn := databases.NewDatabase().DbConnection()
	defer func() {
		if err := dbconn.Close(); err != nil {
			log.Printf("Fail to db connection close. %v\n", err)
		}
	}()

	gin.SetMode(getGinMode())
	r := router.InitRouter(dbconn)

	s := &http.Server{
		Addr:    config.Conf.Web.Host + ":" + config.Conf.Web.Port,
		Handler: r,
	}
	s.ListenAndServe()
}

func getGinMode() string {
	// AppEnv : dev or stg or prd
	appEnv := config.Conf.App.AppEnv
	if appEnv == "prd" {
		return gin.ReleaseMode
	}
	return gin.DebugMode
}
