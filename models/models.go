package models

import (
	"fmt"

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

type CompanyType int

const (
	Corporations CompanyType = iota
	NonProfit
	Cooperative
	SoleProprietorship
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"email;not null;unique"`
	Password string `json:"password" gorm:"password;not null"`
}

type Company struct {
	ID                string      `gorm:"id;primarykey"`
	Name              string      `gorm:"name;not null;unique"`
	Description       string      `gorm:"description"`
	AmountOfEmployees int         `gorm:"amount_of_employees;not null"`
	Registered        bool        `gorm:"registered;not null"`
	Type              CompanyType `gorm:"not null"`
}
