package handlers

import (
	"github.com/jinzhu/gorm"

	"clean-architechure-golang/controllers"
	"clean-architechure-golang/repositories"
	"clean-architechure-golang/services"
)

func MusicHandler(dbconn *gorm.DB) controllers.MusicController {
	musicRepository := repositories.NewMusicRepository(dbconn)
	artistRepository := repositories.NewArtistRepository(dbconn)
	companyRepository := repositories.NewCompanyRepository(dbconn)
	musicService := services.NewMusicService(musicRepository, artistRepository, companyRepository)
	musicController := controllers.NewMusicController(musicService)
	return musicController
}
