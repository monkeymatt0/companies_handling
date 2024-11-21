package services

import (
	"companies_handling/models"
	"companies_handling/repositories"
)

type CompanyService interface {
	CreateCompany(company *models.Company) (*string, error)
	EditCompany(company *models.Company) (*models.Company, error)
	GetCompany(uuid string) (*models.Company, error)
	GetCompanyUser(id uint, uuid string) (*models.Company, error)
	DeleteCompany(uuid string) error
	DeleteCompanyHard(uuid string) error
}

type companyService struct {
	repository repositories.CompanyRepository
}

func (cs *companyService) CreateCompany(company *models.Company) (*string, error) {
	return cs.repository.CreateCompany(company)
}
func (cs *companyService) EditCompany(company *models.Company) (*models.Company, error) {
	return cs.repository.EditCompany(company)
}
func (cs *companyService) GetCompany(uuid string) (*models.Company, error) {
	return cs.repository.GetCompany(uuid)
}
func (cs *companyService) GetCompanyUser(id uint, uuid string) (*models.Company, error) {
	return cs.repository.GetCompanyUser(id, uuid)
}
func (cs *companyService) DeleteCompany(uuid string) error {
	return cs.repository.DeleteCompany(uuid)
}
func (cs *companyService) DeleteCompanyHard(uuid string) error {
	return cs.repository.DeleteCompanyHard(uuid)
}

func NewCompanyService(repository repositories.CompanyRepository) CompanyService {
	return &companyService{repository: repository}
}
