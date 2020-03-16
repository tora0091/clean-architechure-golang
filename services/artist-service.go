package services

import (
	"github.com/gin-gonic/gin"

	"clean-architechure-golang/entities"
	"clean-architechure-golang/repositories"
)

type ArtistService interface {
	FindAll() ([]*entities.Artist, error)
	Save(c *gin.Context) (*entities.Artist, error)
}

type artistService struct {
	repository repositories.ArtistRepository
}

func NewArtistService(repository repositories.ArtistRepository) ArtistService {
	return &artistService{
		repository: repository,
	}
}

func (service *artistService) FindAll() ([]*entities.Artist, error) {
	return service.repository.FindAll()
}

func (service *artistService) Save(c *gin.Context) (*entities.Artist, error) {
	artist, err := getRequestParam(c)
	if err != nil {
		return nil, err
	}

	// TODO: validation

	savedArtist, err := service.repository.Save(artist)
	if err != nil {
		return nil, err
	}
	return savedArtist, nil
}

func getRequestParam(c *gin.Context) (*entities.Artist, error) {
	var artist entities.Artist

	err := c.BindJSON(&artist)
	if err != nil {
		return nil, err
	}
	return &artist, nil
}
