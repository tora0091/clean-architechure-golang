package router

import (
	"clean-architechure-golang/config"
	"clean-architechure-golang/jsons"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"clean-architechure-golang/handlers"
)

func InitRouter(dbconn *gorm.DB) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/musics", handlers.MusicHandler(dbconn).FindAll)
		v1.POST("/music", handlers.MusicHandler(dbconn).Save)
		v1.GET("/music/:id", handlers.MusicHandler(dbconn).FindByID)
		v1.PUT("/music/:id", handlers.MusicHandler(dbconn).UpdateByID)
		v1.DELETE("/music/:id", handlers.MusicHandler(dbconn).DeleteByID)

		v1.GET("/artists", handlers.ArtistHandler(dbconn).FindAll)
		v1.POST("/artist", handlers.ArtistHandler(dbconn).Save)
		v1.GET("/artist/:id", handlers.ArtistHandler(dbconn).FindByID)
		v1.PUT("/artist/:id", handlers.ArtistHandler(dbconn).UpdateByID)
		v1.DELETE("/artist/:id", handlers.ArtistHandler(dbconn).DeleteByID)

		v1.GET("/companies", handlers.CompanyHandler(dbconn).FindAll)
		v1.POST("/company", handlers.CompanyHandler(dbconn).Save)
		v1.GET("/company/:id", handlers.CompanyHandler(dbconn).FindByID)
		v1.PUT("/company/:id", handlers.CompanyHandler(dbconn).UpdateByID)
		v1.DELETE("/company/:id", handlers.CompanyHandler(dbconn).DeleteByID)
	}

	v2 := r.Group("/api/v2")
	{
		v2.GET("/musics", handlers.MusicHandler(dbconn).FindAllData)
		v2.POST("/music", handlers.MusicHandler(dbconn).SaveAllData)
		v2.GET("/music/:id", handlers.MusicHandler(dbconn).FindAllDataByID)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, jsons.ResponseMessage{Message: config.Conf.Message.RouteNotFound})
	})
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, jsons.ResponseMessage{Message: config.Conf.Message.MethodNotAllowed})
	})
	return r
}
