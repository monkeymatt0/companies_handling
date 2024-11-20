package internal

import "gorm.io/gorm"

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

type CompanyType uint8

const (
	Corporations CompanyType = iota
	NonProfit
	Cooperative
	SoleProprietorship
)

type User struct {
	gorm.Model
	Email    string `gorm:"email;not null;unique"`
	Password string `gorm:"password:not null"`
}

type Company struct {
	gorm.Model
	Name              string      `gorm:"name;not null;unique"`
	Description       string      `gorm:"description;"`
	AmountOfEmployees uint        `gorm:"amount_of_employees;not null"`
	Registered        bool        `gorm:"registered;not null"`
	Type              CompanyType `gorm:"type:not null"`
}
