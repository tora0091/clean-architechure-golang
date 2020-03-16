package services

import (
	"time"

	"github.com/gin-gonic/gin"

	"clean-architechure-golang/entities"
	"clean-architechure-golang/repositories"
)

type CompanyService interface {
	FindAll() ([]*entities.Company, error)
	Save(c *gin.Context) (*entities.Company, error)
	FindByID(c *gin.Context) (*entities.Company, error)
	UpdateByID(c *gin.Context) (*entities.Company, error)
	DeleteByID(c *gin.Context) error
}

type companyService struct {
	repository repositories.CompanyRepository
}

func NewCompanyService(repository repositories.CompanyRepository) CompanyService {
	return &companyService{
		repository: repository,
	}
}

func (service *companyService) FindAll() ([]*entities.Company, error) {
	return service.repository.FindAll()
}

func (service *companyService) Save(c *gin.Context) (*entities.Company, error) {
	company, err := getRequestParamCompany(c)
	if err != nil {
		return nil, err
	}

	// TODO: validation

	savedCompany, err := service.repository.Save(company)
	if err != nil {
		return nil, err
	}
	return savedCompany, nil
}

func (service *companyService) FindByID(c *gin.Context) (*entities.Company, error) {
	id, err := getIDParam(c)
	if err != nil {
		return nil, err
	}

	company, err := service.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (service *companyService) UpdateByID(c *gin.Context) (*entities.Company, error) {
	id, err := getIDParam(c)
	if err != nil {
		return nil, err
	}

	updateCompany, err := getRequestParamCompany(c)
	if err != nil {
		return nil, err
	}

	// TODO: validation

	updateCompany.UpdatedAt = time.Now()

	company, err := service.repository.UpdateByID(id, updateCompany)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (service *companyService) DeleteByID(c *gin.Context) error {
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
