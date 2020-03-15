package main

import (
	"github.com/gin-gonic/gin"

	"clean-architechure-golang/controllers"
	"clean-architechure-golang/repositories"
	"clean-architechure-golang/services"
)

var (
	musicService    = services.NewMusicService()
	musicController = controllers.NewMusicController(musicService)

	artistRepository = repositories.NewArtistRepository()
	artistService    = services.NewArtistService(artistRepository)
	artistController = controllers.NewArtistController(artistService)
)

func main() {
	r := gin.Default()

	music := r.Group("/api/v1/music")
	{
		music.GET("/list", musicController.FindAll)
	}

	artist := r.Group("/api/v1/artist")
	{
		artist.GET("/list", artistController.FindAll)
	}

	company := r.Group("/api/v1/company")
	{
		company.GET("/list")
	}

	r.Run(":8080")
}
