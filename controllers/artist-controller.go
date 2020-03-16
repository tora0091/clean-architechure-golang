package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"clean-architechure-golang/jsons"
	"clean-architechure-golang/services"
)

type ArtistController interface {
	FindAll(c *gin.Context)
	Save(c *gin.Context)
}

type artistController struct {
	service services.ArtistService
}

func NewArtistController(service services.ArtistService) ArtistController {
	return &artistController{
		service: service,
	}
}

// curl -v -X GET http://localhost:8080/api/v1/artist/list
func (controller *artistController) FindAll(c *gin.Context) {
	artists, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
	}
	c.JSON(http.StatusOK, jsons.StatusAndArtistList{Status: http.StatusOK, Data: artists})
}

// curl -v -X POST -H "Content-type: application/json" -d '{"name":"ai kimura","email":"izumi@example.com","birth":"19970903","gender":"nomal","address":"Tokyo"}' http://localhost:8080/api/v1/artist
func (controller *artistController) Save(c *gin.Context) {
	artist, err := controller.service.Save(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
	}
	c.JSON(http.StatusOK, jsons.StatusAndArtist{Status: http.StatusOK, Data: artist})
}
