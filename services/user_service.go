package services

import (
	"companies_handling/models"
	"companies_handling/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUser(id uint64) (*models.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func (us *userService) CreateUser(user *models.User) error {
	return us.repository.CreateUser(user)
}

func (us *userService) GetUser(id uint64) (*models.User, error) {
	return us.repository.GetUser(id)
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository: repository}
}
