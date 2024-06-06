package postgres

import (
	"fmt"

	"github.com/Go11Group/Javokhir-A/at_lesson/GormPractice/models"
	"gorm.io/gorm"
)

type NewUserRepository struct {
	Db *gorm.DB
}

func CreateNewUserRepo(db *gorm.DB) *NewUserRepository {

	return &NewUserRepository{Db: db}
}

func (u *NewUserRepository) Create(user models.User) error {
	if err := u.Db.Create(&user).Error; err != nil {
		return fmt.Errorf("failed while creating user %w", err)
	}

	return nil
}

func (u *NewUserRepository) Update(user models.User) error {

	updates := map[string]interface{}{
		"first_name":  user.FirstName,
		"last_name":   user.LastName,
		"Email":       user.Email,
		"Password":    user.Password,
		"age":         user.Age,
		"field":       user.Filed,
		"is_employee": user.IsEmploee,
		"gender":      user.Geneder,
	}

	u.Db.Model(&user).Where("id = ?", user.ID).Updates(updates)

	return nil
}
func (u *NewUserRepository) Delete(user *models.User) error {

	if err := u.Db.Where("id = ?", user.ID).Delete(&user).Error; err != nil {
		return fmt.Errorf("failed while deleting user record %w", err)
	}

	return nil
}

// func (u *NewUserRepository) GetUserById(userId int) (*models.User, error) {
// 	user := models.User{Model.r: userId}
// 	u.Db.Model(userId)
// }
