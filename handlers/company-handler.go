package handlers

import (
	"github.com/jinzhu/gorm"

	"clean-architechure-golang/controllers"
	"clean-architechure-golang/repositories"
	"clean-architechure-golang/services"
)

func CompanyHandler(dbconn *gorm.DB) controllers.CompanyController {
	companyRepository := repositories.NewCompanyRepository(dbconn)
	companyService := services.NewCompanyService(companyRepository)
	companyController := controllers.NewCompanyController(companyService)
	return companyController
}
