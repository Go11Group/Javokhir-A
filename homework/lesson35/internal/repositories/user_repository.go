package repositories

import (
	"database/sql"
	"fmt"

	"github.com/Go11Group/Javokhir-A/homework/lesson35/internal/models"
	"github.com/google/uuid"
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
	id := uuid.NewString()
	query := `
		INSERT INTO USERS(id, name, user_name, password)
		VALUES($1, $2, $3, $4)
	`
	_, err := u.db.Exec(query, id, user.Name, user.UserName, user.Password)
	if err != nil {
		return fmt.Errorf("Creating user failed:" + err.Error())
	}

	return nil
}

func (u UserRepository) GetUser(userID string) (*models.User, error) {
	query := `
		SELECT id, name, user_name, password, rank 
		FROM users
		WHERE deleted_at IS NULL and id = $1;
	`
	row := u.db.QueryRow(query, userID)
	user := models.User{}

	err := row.Scan(&user.Id, &user.Name, &user.UserName, &user.Password, &user.Rank)

	if err != nil {
		return nil, fmt.Errorf("getting user failed:" + err.Error())
	}

	return &user, nil
}

func (u UserRepository) DeleteUser(userId string) {

}
