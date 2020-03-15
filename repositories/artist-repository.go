package repositories

import (
	"clean-architechure-golang/entities"
)

type ArtistRepository interface {
	FindAll() []*entities.Artist
}

type artistRepository struct{}

func NewArtistRepository() ArtistRepository {
	return &artistRepository{}
}

func (repository *artistRepository) FindAll() []*entities.Artist {
	return nil
}
