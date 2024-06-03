package repositories

import (
	"fmt"

	"github.com/Go11Group/Javokhir-A/homework/lesson30/transactions/internal/app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (u *UserRepository) CreateUser(user *models.User) error {
	return u.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return fmt.Errorf("creating user failed: %v", err)
		}
		return nil
	})
}

func (u *UserRepository) UpdateUser(user *models.User) error {
	tx := u.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return fmt.Errorf("beginning transaction failed: %v", err)
	}

	if err := tx.Model(&models.User{}).Where("id = ?", user.ID).Updates(
		map[string]interface{}{
			"user_name": user.UserName,
			"email":     user.Email,
			"password":  user.Password,
		}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("updating user failed: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commiting transaction failed: %v", err)
	}

	return nil
}

func (u *UserRepository) DeleterUser(userId int) error {
	tx := u.Db.Begin()
	if err := tx.Error; err != nil {
		return fmt.Errorf("beginning transaction failed: %v", err)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()

		}
	}()
	var user *models.User

	if err := tx.Model(&models.User{}).Where("id = ?", userId).Find(&user).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("finding user by id to delete failed: %v", err)
	}

	if err := tx.Delete(user).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("deleting user failed: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commiting user transaction failed: %v", err)
	}

	return nil
}

func (u *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User

	err := u.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, id).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}
