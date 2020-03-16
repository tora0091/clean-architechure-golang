package main

import (
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
