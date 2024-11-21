package handlers

import (
	"companies_handling/models"
	"companies_handling/services"
	"companies_handling/utils"
	"companies_handling/validators"
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not valid ID"})
		return
	}
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide the ID"})
		return
	}

	if _, err := utils.RightCompany(uint(id), uuid, h.companyService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with the user ID"})
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not valid ID"})
		return
	}

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
	company.UserID = uint(id)

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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not valid ID"})
		return
	}

	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide the ID"})
		return
	}
	if _, err := utils.RightCompany(uint(id), uuid, h.companyService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with the user ID"})
		return
	}
	var company *models.Company
	company, err2 := h.companyService.GetCompany(uuid)
	if err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, *company)
}

func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not valid ID"})
		return
	}
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide the ID"})
		return
	}
	if _, err := utils.RightCompany(uint(id), uuid, h.companyService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with the user ID"})
		return
	}
	if err := h.companyService.DeleteCompany(uuid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem while deleting"})
		return
	}

	c.JSON(http.StatusOK, uuid)
}

func (h *CompanyHandler) DeleteCompanyHard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not valid ID"})
		return
	}

	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide the ID"})
		return
	}
	if _, err := utils.RightCompany(uint(id), uuid, h.companyService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with the user ID"})
		return
	}
	if err := h.companyService.DeleteCompanyHard(uuid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem while deleting"})
		return
	}

	c.JSON(http.StatusOK, uuid)
}
