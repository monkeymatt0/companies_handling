package services

import (
	"companies_handling/models"
	"companies_handling/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUser(id int) (*models.User, error)
	DeleteUser(id int) error
	DeleteUserHard(id int) error
	GetUserByEmail(email string) (*models.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func (us *userService) CreateUser(user *models.User) error {
	return us.repository.CreateUser(user)
}

func (us *userService) GetUser(id int) (*models.User, error) {
	return us.repository.GetUser(id)
}

func (us *userService) GetUserByEmail(email string) (*models.User, error) {
	return us.repository.GetUserByEmail(email)
}

func (us *userService) DeleteUser(id int) error {
	return us.repository.DeleteUser(id)
}

func (us *userService) DeleteUserHard(id int) error {
	return us.repository.DeleteUserHard(id)
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository: repository}
}
