package repositories

import (
	"github.com/Go11Group/Javokhir-A/homework/lesson33/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	return &user, result.Error
}

func (r *UserRepository) CreateUser(user *models.User) error {
	result := r.db.Create(user)
	return result.Error
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	result := r.db.Save(user)
	return result.Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	return result.Error
}
