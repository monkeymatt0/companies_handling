package repositories

import (
	"companies_handling/models"

	"gorm.io/gorm"
)

// CreateUser is used to create a single user
// GetUser is used to fetch a single user end return it
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(id int) (*models.User, error)
	DeleteUser(id int) error
	DeleteUserHard(id int) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (ur *userRepository) CreateUser(user *models.User) error {
	err := ur.db.Create(&user).Error
	return err
}

func (ur *userRepository) GetUser(id int) (*models.User, error) {
	var user models.User
	err := ur.db.Find(&user, id).Error
	user.Password = ""
	return &user, err
}

func (ur *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := ur.db.Where("email = ?", email).Find(&user).Error
	user.Password = ""
	return &user, err
}

func (ur *userRepository) DeleteUser(id int) error {
	err := ur.db.Delete(&models.User{}, id).Error
	return err
}

func (ur *userRepository) DeleteUserHard(id int) error {
	err := ur.db.Unscoped().Delete(&models.User{}, id).Error
	return err
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
