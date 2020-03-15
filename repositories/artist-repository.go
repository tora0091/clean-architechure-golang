package repositories

import (
	"github.com/jinzhu/gorm"

	"clean-architechure-golang/entities"
)

type ArtistRepository interface {
	FindAll() []*entities.Artist
}

type artistRepository struct {
	conn *gorm.DB
}

func NewArtistRepository(conn *gorm.DB) ArtistRepository {
	return &artistRepository{
		conn: conn,
	}
}

func (repository *artistRepository) FindAll() []*entities.Artist {
	var artist entities.Artist
	repository.conn.Table("artist").Find(&artist)
	return nil
}
