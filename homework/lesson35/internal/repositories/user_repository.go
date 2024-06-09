package repositories

import (
	"database/sql"
	"fmt"

	"github.com/Go11Group/Javokhir-A/homework/lesson35/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) CreateUser(user *models.User) error {
	fmt.Println("User repo")
	return nil
}
