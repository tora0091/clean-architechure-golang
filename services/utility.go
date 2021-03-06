package services

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"clean-architechure-golang/entities"
)

func getRequestParamMusic(c *gin.Context) (*entities.Music, error) {
	var music entities.Music

	err := c.BindJSON(&music)
	if err != nil {
		return nil, err
	}
	return &music, nil
}

func getRequestParamArtist(c *gin.Context) (*entities.Artist, error) {
	var artist entities.Artist

	err := c.BindJSON(&artist)
	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func getRequestParamCompany(c *gin.Context) (*entities.Company, error) {
	var company entities.Company

	err := c.BindJSON(&company)
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func getRequestParamMusicAll(c *gin.Context) (*entities.MusicStructResponse, error) {
	var music entities.MusicStructResponse

	err := c.BindJSON(&music)
	if err != nil {
		return nil, err
	}
	return &music, nil
}

func getIDParam(c *gin.Context) (int, error) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return id, err
	}
	return id, nil
}
