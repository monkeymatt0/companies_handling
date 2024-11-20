package repositories

import (
	"companies_handling/models"

	"gorm.io/gorm"
)

// CreateUser is used to create a single user
// GetUser is used to fetch a single user end return it
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(id uint64) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (ur *userRepository) CreateUser(user *models.User) error {
	err := ur.db.Create(&user).Error
	return err
}

func (ur *userRepository) GetUser(id uint64) (models.User, error) {
	var user models.User
	err := ur.db.Find(&user, id).Error
	return user, err
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
