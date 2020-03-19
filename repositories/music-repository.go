package repositories

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"clean-architechure-golang/config"
	"clean-architechure-golang/entities"
)

type MusicRepository interface {
	FindAll() ([]*entities.Music, error)
	Save(*entities.Music) (*entities.Music, error)
	FindByID(int) (*entities.Music, error)
	UpdateByID(int, *entities.Music) (*entities.Music, error)
	DeleteByID(int) error
}

type musicRepository struct {
	conn *gorm.DB
}

func NewMusicRepository(conn *gorm.DB) MusicRepository {
	return &musicRepository{
		conn: conn,
	}
}

func (repository *musicRepository) FindAll() ([]*entities.Music, error) {
	var musics []*entities.Music
	repository.conn.Table(config.Conf.Table.Music).Find(&musics)
	if len(musics) > 0 {
		return musics, nil
	}
	return nil, fmt.Errorf(config.Conf.Message.RecordNotFound)
}

func (repository *musicRepository) Save(music *entities.Music) (*entities.Music, error) {
	db := repository.conn.Table(config.Conf.Table.Music)

	db.NewRecord(music)
	db.Create(music)
	if db.NewRecord(music) == false {
		return music, nil
	}
	return nil, fmt.Errorf(config.Conf.Message.SaveError)
}

func (repository *musicRepository) FindByID(id int) (*entities.Music, error) {
	db := repository.conn.Table(config.Conf.Table.Music)

	var music entities.Music
	recordNotFound := db.First(&music, id).RecordNotFound()
	if recordNotFound {
		return nil, fmt.Errorf(config.Conf.Message.RecordNotFound)
	}
	return &music, nil
}

func (repository *musicRepository) UpdateByID(id int, updateData *entities.Music) (*entities.Music, error) {
	db := repository.conn.Table(config.Conf.Table.Music)

	var music entities.Music
	if err := db.First(&music, id).Update(updateData).Error; err != nil {
		return nil, err
	}
	return &music, nil
}

func (repository *musicRepository) DeleteByID(id int) error {
	db := repository.conn.Table(config.Conf.Table.Music)

	var music entities.Music
	if err := db.First(&music, id).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}
