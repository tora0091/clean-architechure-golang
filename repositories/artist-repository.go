package repositories

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"clean-architechure-golang/config"
	"clean-architechure-golang/entities"
)

type ArtistRepository interface {
	FindAll() ([]*entities.Artist, error)
	Save(*entities.Artist) (*entities.Artist, error)
}

type artistRepository struct {
	conn *gorm.DB
}

func NewArtistRepository(conn *gorm.DB) ArtistRepository {
	return &artistRepository{
		conn: conn,
	}
}

func (repository *artistRepository) FindAll() ([]*entities.Artist, error) {
	var artists []*entities.Artist
	repository.conn.Table(config.Conf.Table.Artist).Find(&artists)
	if len(artists) > 0 {
		return artists, nil
	}
	return nil, fmt.Errorf("Record Not Found")
}

func (repository *artistRepository) Save(artist *entities.Artist) (*entities.Artist, error) {
	db := repository.conn.Table(config.Conf.Table.Artist)
	db.NewRecord(artist)
	db.Create(artist)
	if db.NewRecord(artist) == false {
		return artist, nil
	}
	return nil, fmt.Errorf("New Record Save Error")
}
