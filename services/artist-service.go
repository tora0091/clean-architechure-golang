package services

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"clean-architechure-golang/entities"
	"clean-architechure-golang/repositories"
)

type ArtistService interface {
	FindAll() ([]*entities.Artist, error)
	Save(c *gin.Context) (*entities.Artist, error)
	FindByID(c *gin.Context) (*entities.Artist, error)
	UpdateByID(c *gin.Context) (*entities.Artist, error)
	DeleteByID(c *gin.Context) error
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

func (service *artistService) FindByID(c *gin.Context) (*entities.Artist, error) {
	id, err := getIDParam(c)
	if err != nil {
		return nil, err
	}

	artist, err := service.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return artist, nil
}

func (service *artistService) UpdateByID(c *gin.Context) (*entities.Artist, error) {
	id, err := getIDParam(c)
	if err != nil {
		return nil, err
	}

	updateArtist, err := getRequestParam(c)
	if err != nil {
		return nil, err
	}

	// TODO: validation

	updateArtist.UpdatedAt = time.Now()

	artist, err := service.repository.UpdateByID(id, updateArtist)
	if err != nil {
		return nil, err
	}
	return artist, nil
}

func (service *artistService) DeleteByID(c *gin.Context) error {
	id, err := getIDParam(c)
	if err != nil {
		return err
	}

	err = service.repository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}

func getRequestParam(c *gin.Context) (*entities.Artist, error) {
	var artist entities.Artist

	err := c.BindJSON(&artist)
	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func getIDParam(c *gin.Context) (int, error) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return id, err
	}
	return id, nil
}
