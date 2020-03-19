package repositories

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"clean-architechure-golang/config"
	"clean-architechure-golang/entities"
)

type CompanyRepository interface {
	FindAll() ([]*entities.Company, error)
	Save(*entities.Company) (*entities.Company, error)
	FindByID(int) (*entities.Company, error)
	UpdateByID(int, *entities.Company) (*entities.Company, error)
	DeleteByID(int) error
}

type companyRepository struct {
	conn *gorm.DB
}

func NewCompanyRepository(conn *gorm.DB) CompanyRepository {
	return &companyRepository{
		conn: conn,
	}
}

func (repository *companyRepository) FindAll() ([]*entities.Company, error) {
	var company []*entities.Company
	repository.conn.Table(config.Conf.Table.Company).Find(&company)
	if len(company) > 0 {
		return company, nil
	}
	return nil, fmt.Errorf(config.Conf.Message.RecordNotFound)
}

func (repository *companyRepository) Save(company *entities.Company) (*entities.Company, error) {
	db := repository.conn.Table(config.Conf.Table.Company)

	db.NewRecord(company)
	db.Create(company)
	if db.NewRecord(company) == false {
		return company, nil
	}
	return nil, fmt.Errorf(config.Conf.Message.SaveError)
}

func (repository *companyRepository) FindByID(id int) (*entities.Company, error) {
	db := repository.conn.Table(config.Conf.Table.Company)

	var company entities.Company
	recordNotFound := db.First(&company, id).RecordNotFound()
	if recordNotFound {
		return nil, fmt.Errorf(config.Conf.Message.RecordNotFound)
	}
	return &company, nil
}

func (repository *companyRepository) UpdateByID(id int, updateData *entities.Company) (*entities.Company, error) {
	db := repository.conn.Table(config.Conf.Table.Company)

	var company entities.Company
	if err := db.First(&company, id).Update(updateData).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (repository *companyRepository) DeleteByID(id int) error {
	db := repository.conn.Table(config.Conf.Table.Company)

	var company entities.Company
	if err := db.First(&company, id).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}
