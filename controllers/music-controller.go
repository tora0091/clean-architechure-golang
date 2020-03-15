package controllers

import (
	"github.com/gin-gonic/gin"

	"clean-architechure-golang/services"
)

type MusicController interface {
	FindAll(*gin.Context)
}

type musicController struct {
	service services.MusicService
}

func NewMusicController(service services.MusicService) MusicController {
	return &musicController{
		service: service,
	}
}

func (controller *musicController) FindAll(c *gin.Context) {
	controller.service.FindAll()
}
