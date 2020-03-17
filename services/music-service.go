package services

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"clean-architechure-golang/entities"
	"clean-architechure-golang/repositories"
)

type MusicService interface {
	FindAll() ([]*entities.Music, error)
	Save(c *gin.Context) (*entities.Music, error)
	FindByID(c *gin.Context) (*entities.Music, error)
	UpdateByID(c *gin.Context) (*entities.Music, error)
	DeleteByID(c *gin.Context) error
	FindAllData() ([]*entities.MusicStructResponse, error)
	SaveAllData(c *gin.Context) (*entities.MusicStructResponse, error)
	FindAllDataByID(c *gin.Context) (*entities.MusicStructResponse, error)
}

type musicService struct {
	conn              *gorm.DB
	repository        repositories.MusicRepository
	artistRepository  repositories.ArtistRepository
	companyRepository repositories.CompanyRepository
}

func NewMusicService(conn *gorm.DB, repository repositories.MusicRepository, artistRepository repositories.ArtistRepository, companyRepository repositories.CompanyRepository) MusicService {
	return &musicService{
		conn:              conn,
		repository:        repository,
		artistRepository:  artistRepository,
		companyRepository: companyRepository,
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

	if err = service.checkIDs(music); err != nil {
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

	if err = service.checkIDs(updateMusic); err != nil {
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

func (service *musicService) FindAllData() ([]*entities.MusicStructResponse, error) {
	musics, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var musicResp []*entities.MusicStructResponse
	for _, music := range musics {
		artist, _ := service.getArtistByID(music)
		company, _ := service.getCompanyByID(music)
		resp := entities.MusicStructResponse{
			ID:        music.ID,
			ISWC:      music.ISWC,
			Title:     music.Title,
			Time:      music.Time,
			Genre:     music.Genre,
			Artist:    artist,
			Company:   company,
			CreatedAt: music.CreatedAt,
			UpdatedAt: music.UpdatedAt,
			DeletedAt: music.DeletedAt,
		}
		musicResp = append(musicResp, &resp)
	}
	return musicResp, nil
}

func (service *musicService) SaveAllData(c *gin.Context) (*entities.MusicStructResponse, error) {
	content, err := getRequestParamMusicAll(c)
	if err != nil {
		return nil, err
	}

	tx := service.conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Error; err != nil {
		return nil, err
	}

	// Wow!
	artist, err := repositories.NewArtistRepository(tx).Save(content.Artist)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Artist Data is Save Error")
	}

	// oh no...
	company, err := repositories.NewCompanyRepository(tx).Save(content.Company)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Company Data is Save Error")
	}

	resp := entities.Music{
		ID:        content.ID,
		ISWC:      content.ISWC,
		Title:     content.Title,
		Time:      content.Time,
		Genre:     content.Genre,
		ArtistID:  artist.ID,
		CompanyID: company.ID,
		CreatedAt: content.CreatedAt,
		UpdatedAt: content.UpdatedAt,
		DeletedAt: content.DeletedAt,
	}

	// i dont know... help me...
	music, err := repositories.NewMusicRepository(tx).Save(&resp)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Music Data is Save Error")
	}
	tx.Commit()

	content.ID = music.ID
	return content, tx.Error
}

func (service *musicService) FindAllDataByID(c *gin.Context) (*entities.MusicStructResponse, error) {
	id, err := getIDParam(c)
	if err != nil {
		return nil, err
	}

	music, err := service.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	artist, _ := service.getArtistByID(music)
	company, _ := service.getCompanyByID(music)
	resp := entities.MusicStructResponse{
		ID:        music.ID,
		ISWC:      music.ISWC,
		Title:     music.Title,
		Time:      music.Time,
		Genre:     music.Genre,
		Artist:    artist,
		Company:   company,
		CreatedAt: music.CreatedAt,
		UpdatedAt: music.UpdatedAt,
		DeletedAt: music.DeletedAt,
	}
	return &resp, nil
}

func (service *musicService) checkIDs(music *entities.Music) error {
	_, err := service.getArtistByID(music)
	if err != nil {
		return err
	}

	_, err = service.getCompanyByID(music)
	if err != nil {
		return err
	}
	return nil
}

func (service *musicService) getArtistByID(music *entities.Music) (*entities.Artist, error) {
	artist, err := service.artistRepository.FindByID(music.ArtistID)
	if err != nil {
		return nil, fmt.Errorf("Artist ID is not Found. Check your artist ID")
	}
	return artist, nil
}

func (service *musicService) getCompanyByID(music *entities.Music) (*entities.Company, error) {
	company, err := service.companyRepository.FindByID(music.CompanyID)
	if err != nil {
		return nil, fmt.Errorf("Company ID is not Found. Check your company ID")
	}
	return company, nil
}
