package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Go11Group/Javokhir-A/at_lesson/lesson29Practice/models"
	"github.com/google/uuid"
)

type NewUsersRepository struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *NewUsersRepository {
	return &NewUsersRepository{Db: db}
}

func (u *NewUsersRepository) GetAllUsers() ([]models.User, error) {
	query := `
	SELECT user_id, first_name, last_name, email, gender, age 
	FROM USERS
	`
	users := []models.User{}

	rows, err := u.Db.Query(query)
	if err != nil {
		return []models.User{}, fmt.Errorf("failed while quering GetAllUsers: %w", err)
	}
	defer rows.Close()

	user := models.User{}
	for rows.Next() {
		err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Gender, &user.Age)
		if err != nil {
			return nil, fmt.Errorf("failed while scaning %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed while iterating over rows error: %w", err)
	}

	return users, nil
}

func (u *NewUsersRepository) GetUserById(userId string) (*models.User, error) {
	query := `
		SELECT user_id, first_name, last_name, email, gender, age 
		FROM USERS
		WHERE user_id = $1
	`
	user := models.User{}

	row := u.Db.QueryRow(query, userId)

	err := row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Gender, &user.Age)
	if err != nil {
		return nil, fmt.Errorf("failed while scanning into user struct %w", err)
	}

	return &user, nil
}

func (u *NewUsersRepository) Create(user models.User) error {
	query := `
		INSERT INTO users(user_id, first_name, last_name, email, gender, age) 
		VALUES($1, $2, $3, $4, $5, $6)
	`

	_, err := u.Db.Exec(query, uuid.NewString(), user.FirstName, user.LastName, user.Email, user.Gender, user.Age)

	if err != nil {
		return fmt.Errorf("failed while inserting user into database %w", err)
	}

	return nil
}
