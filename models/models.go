package models

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// @remind : Develop models here
type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		c.Database.Host,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
		c.Database.Port,
	)
}

type CompanyType string

const (
	Corporation        CompanyType = "Corporation"
	NonProfit          CompanyType = "NonProfit"
	Cooperative        CompanyType = "Cooperative"
	SoleProprietorship CompanyType = "SoleProprietorship"
)

type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"email;not null;unique"`
	Password  string `json:"password" gorm:"password;not null"`
	Companies []Company
}

type Company struct {
	ID                string      `json:"id" gorm:"id;primarykey" validate:"required,uuid4"`
	Name              string      `json:"name" gorm:"name;not null;unique" validate:"required,max=15"`
	Description       string      `json:"description" gorm:"description" validate:"omitempty,max=3000"`
	AmountOfEmployees int         `json:"amountOfEmployees" gorm:"amount_of_employees;not null" validate:"required,min=1"`
	Registered        bool        `json:"registered" gorm:"registered;not null" validate:"required"`
	Type              CompanyType `json:"type" gorm:"not null" validate:"oneof=Corporations NonProfit Cooperative SoleProprietorship"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
	UserID            uint
}

type Claims struct {
	Email string `json:"username"`
	jwt.RegisteredClaims
}
