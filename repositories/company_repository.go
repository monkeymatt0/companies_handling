package repositories

import (
	"companies_handling/models"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	CreateCompany(company *models.Company) (*string, error)
	EditCompany(company *models.Company) (*models.Company, error)
	GetCompany(uuid string) (*models.Company, error)
	DeleteCompany(uuid string) error
}

type companyRepository struct {
	db *gorm.DB
}

func (cr *companyRepository) CreateCompany(company *models.Company) (*string, error) {
	err := cr.db.Create(&company).Error
	if err != nil {
		return nil, err
	}
	return &company.ID, err
}

func (cr *companyRepository) EditCompany(company *models.Company) (*models.Company, error) {
	err := cr.db.Model(&company).Updates(company).Error
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (cr *companyRepository) GetCompany(uuid string) (*models.Company, error) {
	var company models.Company
	err := cr.db.Find(&company, uuid).Error
	if err != nil {
		return nil, err
	}
	return &company, err
}

func (cr *companyRepository) DeleteCompany(uuid string) error {
	err := cr.db.Delete(&models.Company{}, uuid).Error
	return err
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{db: db}
}
