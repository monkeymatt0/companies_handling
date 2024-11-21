package validators

import (
	"companies_handling/models"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateCompany(company *models.Company) error {
	validate := validator.New()

	err := validate.Struct(company)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Validation error for field: '%s': '%s'\n", err.Field(), err.Error())
		}
		return err
	}
	return nil
}
