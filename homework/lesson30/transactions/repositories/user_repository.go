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

	err := u.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user); err != nil {
			return fmt.Errorf("creating user rolled back %v", err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
