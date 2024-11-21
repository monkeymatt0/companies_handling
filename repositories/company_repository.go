package repositories

import (
	"companies_handling/models"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	CreateCompany(company *models.Company) (*string, error)
	EditCompany(company *models.Company) (*models.Company, error)
	GetCompany(uuid string) (*models.Company, error)
	GetCompanyUser(id uint, uuid string) (*models.Company, error)
	DeleteCompany(uuid string) error
	DeleteCompanyHard(uuid string) error
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
	uuid := company.ID
	company.ID = ""
	err := cr.db.Unscoped().Where("id = ?", uuid).Updates(company).Error
	if err != nil {
		return nil, err
	}
	return cr.GetCompany(uuid)
}

func (cr *companyRepository) GetCompany(uuid string) (*models.Company, error) {
	var company models.Company
	err := cr.db.Unscoped().Where("id = ?", uuid).First(&company).Error
	if err != nil {
		return nil, err
	}
	return &company, err
}

func (cr *companyRepository) GetCompanyUser(id uint, uuid string) (*models.Company, error) {
	var company models.Company
	err := cr.db.Unscoped().Where("id = ? AND user_id = ?", uuid, id).First(&company).Error
	if err != nil {
		return nil, err
	}
	return &company, err
}

func (cr *companyRepository) DeleteCompany(uuid string) error {
	err := cr.db.Where("id = ?", uuid).Delete(&models.Company{}).Error
	return err
}

func (cr *companyRepository) DeleteCompanyHard(uuid string) error {
	err := cr.db.Unscoped().Where("id = ?", uuid).Delete(&models.Company{}).Error
	return err
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{db: db}
}
