package handlers

import (
	"github.com/jinzhu/gorm"

	"clean-architechure-golang/controllers"
	"clean-architechure-golang/repositories"
	"clean-architechure-golang/services"
)

func ArtistHandler(dbconn *gorm.DB) controllers.ArtistController {
	artistRepository := repositories.NewArtistRepository(dbconn)
	artistService := services.NewArtistService(artistRepository)
	artistController := controllers.NewArtistController(artistService)
	return artistController
}
