package utils

import (
	"companies_handling/services"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Password checker
func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Check if the user can operate on that specific company
func RightCompany(id uint, uuid string, companyService services.CompanyService) (bool, error) {
	c, err := companyService.GetCompanyUser(id, uuid)
	if err != nil {
		return false, err
	}
	return c != nil, nil
}
