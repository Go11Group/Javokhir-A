package repositories

import (
	"github.com/Go11Group/Javokhir-A/homework/lesson33/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
