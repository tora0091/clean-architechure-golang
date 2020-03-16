package repositories

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"clean-architechure-golang/config"
	"clean-architechure-golang/entities"
)

type ArtistRepository interface {
	FindAll() ([]*entities.Artist, error)
	Save(artist *entities.Artist) (*entities.Artist, error)
	FindByID(id int) (*entities.Artist, error)
	UpdateByID(id int, updateData *entities.Artist) (*entities.Artist, error)
	DeleteByID(id int) error
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

func (repository *artistRepository) FindByID(id int) (*entities.Artist, error) {
	db := repository.conn.Table(config.Conf.Table.Artist)

	var artist entities.Artist
	recordNotFound := db.First(&artist, id).RecordNotFound()
	if recordNotFound {
		return nil, fmt.Errorf("Record Not Found")
	}
	return &artist, nil
}

func (repository *artistRepository) UpdateByID(id int, updateData *entities.Artist) (*entities.Artist, error) {
	db := repository.conn.Table(config.Conf.Table.Artist)

	var artist entities.Artist
	if err := db.First(&artist, id).Update(updateData).Error; err != nil {
		return nil, err
	}
	return &artist, nil
}

func (repository *artistRepository) DeleteByID(id int) error {
	db := repository.conn.Table(config.Conf.Table.Artist)

	var artist entities.Artist
	if err := db.First(&artist, id).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}
