package controllers

import (
	"github.com/gin-gonic/gin"

	"clean-architechure-golang/services"
)

type ArtistController interface {
	FindAll(c *gin.Context)
}

type artistController struct {
	service services.ArtistService
}

func NewArtistController(service services.ArtistService) ArtistController {
	return &artistController{
		service: service,
	}
}

func (controller *artistController) FindAll(c *gin.Context) {
	controller.service.FindAll()
}
