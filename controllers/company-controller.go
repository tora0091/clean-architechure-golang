package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"clean-architechure-golang/jsons"
	"clean-architechure-golang/services"
)

type CompanyController interface {
	FindAll(c *gin.Context)
	Save(c *gin.Context)
	FindByID(c *gin.Context)
	UpdateByID(c *gin.Context)
	DeleteByID(c *gin.Context)
}

type companyController struct {
	service services.CompanyService
}

func NewCompanyController(service services.CompanyService) CompanyController {
	return &companyController{
		service: service,
	}
}

// curl -v -X GET http://localhost:8080/api/v1/companies
func (controller *companyController) FindAll(c *gin.Context) {
	companies, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.StatusAndCompanyList{Status: http.StatusOK, Data: companies})
}

// curl -v -X POST -H "Content-type: application/json" -d '{"name":"Hajime co.,ltd","email":"hajime@example.com","address":"Tokyo in Japan"}' http://localhost:8080/api/v1/company
func (controller *companyController) Save(c *gin.Context) {
	company, err := controller.service.Save(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.StatusAndCompany{Status: http.StatusOK, Data: company})
}

// curl -v -X GET http://localhost:8080/api/v1/company/1
func (controller *companyController) FindByID(c *gin.Context) {
	company, err := controller.service.FindByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.StatusAndCompany{Status: http.StatusOK, Data: company})
}

// curl -v -X PUT -H "Content-type: application/json" -d '{"name":"Maido co.,ltd","email":"maido@example.com","address":"Osaka in Japan"}' http://localhost:8080/api/v1/company/2
func (controller *companyController) UpdateByID(c *gin.Context) {
	company, err := controller.service.UpdateByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.StatusAndCompany{Status: http.StatusOK, Data: company})
}

// curl -v -X DELETE http://localhost:8080/api/v1/company/3
func (controller *companyController) DeleteByID(c *gin.Context) {
	err := controller.service.DeleteByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.StatusAndMessage{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.StatusAndMessage{Status: http.StatusOK, Message: http.StatusText(http.StatusOK)})
}
