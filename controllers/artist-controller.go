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
	FindByID(c *gin.Context)
	UpdateByID(c *gin.Context)
	DeleteByID(c *gin.Context)
}

type artistController struct {
	service services.ArtistService
}

func NewArtistController(service services.ArtistService) ArtistController {
	return &artistController{
		service: service,
	}
}

// curl -v -X GET http://localhost:8080/api/v1/artists
func (controller *artistController) FindAll(c *gin.Context) {
	artists, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.ResponseMessage{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.ResponseArtistList{Data: artists})
}

// curl -v -X POST -H "Content-type: application/json" -d '{"name":"ai kimura","email":"ai@example.com","birth":"19970903","gender":"nomal","address":"Tokyo"}' http://localhost:8080/api/v1/artist
func (controller *artistController) Save(c *gin.Context) {
	artist, err := controller.service.Save(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.ResponseMessage{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.ResponseArtist{Data: artist})
}

// curl -v -X GET http://localhost:8080/api/v1/artist/1
func (controller *artistController) FindByID(c *gin.Context) {
	artist, err := controller.service.FindByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.ResponseMessage{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.ResponseArtist{Data: artist})
}

// curl -v -X PUT -H "Content-type: application/json" -d '{"name":"rei izumi","email":"rei@example.com","birth":"20000903","gender":"boy","address":"Osaka"}' http://localhost:8080/api/v1/artist/2
func (controller *artistController) UpdateByID(c *gin.Context) {
	artist, err := controller.service.UpdateByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.ResponseMessage{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.ResponseArtist{Data: artist})
}

// curl -v -X DELETE http://localhost:8080/api/v1/artist/3
func (controller *artistController) DeleteByID(c *gin.Context) {
	err := controller.service.DeleteByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.ResponseMessage{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.ResponseMessage{Message: http.StatusText(http.StatusOK)})
}
