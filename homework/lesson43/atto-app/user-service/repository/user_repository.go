package repository

import (
	"database/sql"
	"fmt"
)

type UserRepository struct {
	Db *sql.DB
}

type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type UserData struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (u *UserRepository) CreateUser(user CreateUser, userId string) error {
	statement, err := u.Db.Prepare(`
        INSERT INTO users(user_id, first_name, last_name, email, age, phone, password_hash) 
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `)

	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer statement.Close()

	_, err = statement.Exec(userId, user.FirstName, user.LastName, user.Email, user.Age, user.Phone, user.Password)
	if err != nil {
		return fmt.Errorf("error executing statement: %v", err)
	}

	return nil
}

func (u *UserRepository) GetUser(userID string) (*UserData, error) {
	var user UserData
	user.UserID = userID

	query := `
		SELECT 
		user_id, 
		first_name, 
		last_name, 
		email, 
		age, 
		phone, 
		created_at, 
		updated_at 
		FROM users WHERE delete_at IS NULL AND user_id = $1
	`

	if err := u.Db.QueryRow(query, userID).Scan(
		&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Age,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt); err != nil {
		return nil, fmt.Errorf("scanning user data failed: " + err.Error())
	}

	return &user, nil
}
