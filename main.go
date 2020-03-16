package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"clean-architechure-golang/config"
	"clean-architechure-golang/databases"
	"clean-architechure-golang/handlers"
)

// var (
// 	musicService    = services.NewMusicService()
// 	musicController = controllers.NewMusicController(musicService)
// )

func main() {
	config.ReadConf()

	dbconn := databases.NewDatabase().DbConnection()

	r := gin.Default()
	gin.SetMode(getGinMode())

	// music := r.Group("/api/v1/music")
	// {
	// 	music.GET("/list", musicController.FindAll)
	// }

	artist := r.Group("/api/v1")
	{
		artist.GET("/artists", handlers.ArtistHandler(dbconn).FindAll)
		artist.POST("/artist", handlers.ArtistHandler(dbconn).Save)
	}

	// company := r.Group("/api/v1/company")
	// {
	// 	company.GET("/list")
	// }

	r.Run(config.Conf.Web.Host + ":" + config.Conf.Web.Port)
}

func getGinMode() string {
	// AppEnv : dev or stg or prd
	appEnv := config.Conf.App.AppEnv
	if appEnv == "prd" {
		return gin.ReleaseMode
	}
	return gin.DebugMode
}
