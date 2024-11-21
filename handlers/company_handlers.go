package handlers

import (
	"companies_handling/models"
	"companies_handling/services"
	"companies_handling/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type CompanyHandler struct {
	companyService services.CompanyService
}

func NewCompanyHandler(companyService services.CompanyService) *CompanyHandler {
	return &CompanyHandler{companyService: companyService}
}

func (h *CompanyHandler) EditCompany(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide the ID"})
		return
	}
	var company models.Company
	if err := c.ShouldBindBodyWithJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company.ID = uuid
	if err := validators.ValidateCompany(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uCompany, err := h.companyService.EditCompany(&company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, uCompany)
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindBodyWithJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUuid, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	company.ID = newUuid.String()

	if err := validators.ValidateCompany(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uuid, err := h.companyService.CreateCompany(&company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, uuid)
}

func (h *CompanyHandler) GetCompany(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide the ID"})
		return
	}
	var company *models.Company
	company, err := h.companyService.GetCompany(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, *company)
}

func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide the ID"})
		return
	}
	if err := h.companyService.DeleteCompany(uuid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem while deleting"})
		return
	}

	c.JSON(http.StatusOK, uuid)
}

func (h *CompanyHandler) DeleteCompanyHard(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide the ID"})
		return
	}
	if err := h.companyService.DeleteCompanyHard(uuid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem while deleting"})
		return
	}

	c.JSON(http.StatusOK, uuid)
}
