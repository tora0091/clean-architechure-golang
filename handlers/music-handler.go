package handlers

import (
	"github.com/jinzhu/gorm"

	"clean-architechure-golang/controllers"
	"clean-architechure-golang/repositories"
	"clean-architechure-golang/services"
)

func MusicHandler(dbconn *gorm.DB) controllers.MusicController {
	musicRepository := repositories.NewMusicRepository(dbconn)
	musicService := services.NewMusicService(musicRepository)
	musicController := controllers.NewMusicController(musicService)
	return musicController
}
