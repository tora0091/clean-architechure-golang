package services

import (
	"time"

	"github.com/gin-gonic/gin"

	"clean-architechure-golang/entities"
	"clean-architechure-golang/repositories"
)

type MusicService interface {
	FindAll() ([]*entities.Music, error)
	Save(c *gin.Context) (*entities.Music, error)
	FindByID(c *gin.Context) (*entities.Music, error)
	UpdateByID(c *gin.Context) (*entities.Music, error)
	DeleteByID(c *gin.Context) error
}

type musicService struct {
	repository repositories.MusicRepository
}

func NewMusicService(repository repositories.MusicRepository) MusicService {
	return &musicService{
		repository: repository,
	}
}

func (service *musicService) FindAll() ([]*entities.Music, error) {
	return service.repository.FindAll()
}

func (service *musicService) Save(c *gin.Context) (*entities.Music, error) {
	music, err := getRequestParamMusic(c)
	if err != nil {
		return nil, err
	}

	// TODO: validation

	savedMusic, err := service.repository.Save(music)
	if err != nil {
		return nil, err
	}
	return savedMusic, nil
}

func (service *musicService) FindByID(c *gin.Context) (*entities.Music, error) {
	id, err := getIDParam(c)
	if err != nil {
		return nil, err
	}

	music, err := service.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return music, nil
}

func (service *musicService) UpdateByID(c *gin.Context) (*entities.Music, error) {
	id, err := getIDParam(c)
	if err != nil {
		return nil, err
	}

	updateMusic, err := getRequestParamMusic(c)
	if err != nil {
		return nil, err
	}

	// TODO: validation

	updateMusic.UpdatedAt = time.Now()

	music, err := service.repository.UpdateByID(id, updateMusic)
	if err != nil {
		return nil, err
	}
	return music, nil
}

func (service *musicService) DeleteByID(c *gin.Context) error {
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
