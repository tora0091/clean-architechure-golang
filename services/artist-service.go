package services

import (
	"clean-architechure-golang/entities"
	"clean-architechure-golang/repositories"
)

type ArtistService interface {
	FindAll() []*entities.Artist
}

type artistService struct {
	repository repositories.ArtistRepository
}

func NewArtistService(repository repositories.ArtistRepository) ArtistService {
	return &artistService{
		repository: repository,
	}
}

func (service *artistService) FindAll() []*entities.Artist {
	return service.repository.FindAll()
}
