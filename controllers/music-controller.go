package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"clean-architechure-golang/jsons"
	"clean-architechure-golang/services"
)

type MusicController interface {
	FindAll(c *gin.Context)
	Save(c *gin.Context)
	FindByID(c *gin.Context)
	UpdateByID(c *gin.Context)
	DeleteByID(c *gin.Context)
	FindAllData(c *gin.Context)
	SaveAllData(c *gin.Context)
	FindAllDataByID(c *gin.Context)
}

type musicController struct {
	service services.MusicService
}

func NewMusicController(service services.MusicService) MusicController {
	return &musicController{
		service: service,
	}
}

// curl -v -X GET http://localhost:8080/api/v1/musics
func (controller *musicController) FindAll(c *gin.Context) {
	musics, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.StatusAndMusicList{Status: http.StatusOK, Data: musics})
}

// curl -v -X POST -H "Content-type: application/json" -d '{"iswc":"A000111222T","title":"I love you","time":5,"genre":"J-POP","artist_id":1,"company_id":1}' http://localhost:8080/api/v1/music
func (controller *musicController) Save(c *gin.Context) {
	music, err := controller.service.Save(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.StatusAndMusic{Status: http.StatusOK, Data: music})
}

// curl -v -X GET http://localhost:8080/api/v1/music/1
func (controller *musicController) FindByID(c *gin.Context) {
	music, err := controller.service.FindByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.StatusAndMusic{Status: http.StatusOK, Data: music})
}

// curl -v -X PUT -H "Content-type: application/json" -d '{"iswc":"A999999999T","title":"why dont you love me","time":7,"genre":"J-POP","artist_id":3,"company_id":2}' http://localhost:8080/api/v1/music/2
func (controller *musicController) UpdateByID(c *gin.Context) {
	music, err := controller.service.UpdateByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.StatusAndMusic{Status: http.StatusOK, Data: music})
}

// curl -v -X DELETE http://localhost:8080/api/v1/music/3
func (controller *musicController) DeleteByID(c *gin.Context) {
	err := controller.service.DeleteByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.StatusAndMessage{Status: http.StatusOK, Message: http.StatusText(http.StatusOK)})
}

// curl -v -X GET http://localhost:8080/api/v2/musics
func (controller *musicController) FindAllData(c *gin.Context) {
	musics, err := controller.service.FindAllData()
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.StatusAndMusicListStructResponse{Status: http.StatusOK, Data: musics})
}

// curl -v -X POST -H "Content-type: application/json" -d '{"iswc":"A123456789T","title":"Are you seeing anyone?","time":9,"genre":"Rock","artist":{"name":"new wave","email":"nami@example.com","birth":"20200101","gender":"mixed","address":"Kyoto"},"company":{"name":"Mad co.,ltd","email":"m@example.com","address":"Kyoto"}}' http://localhost:8080/api/v2/music
func (controller *musicController) SaveAllData(c *gin.Context) {
	music, err := controller.service.SaveAllData(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.StatusAndMusicStructResponse{Status: http.StatusOK, Data: music})
}

// curl -v -X GET http://localhost:8080/api/v2/music/1
func (controller *musicController) FindAllDataByID(c *gin.Context) {
	music, err := controller.service.FindAllDataByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.StatusAndMusicStructResponse{Status: http.StatusOK, Data: music})
}
