package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/models"
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

type UpdateUser struct {
	Name     *string    `json:"name"`
	Email    *string    `json:"user_name"`
	Birthday *time.Time `json:"birthday"`
	Password *string    `json:"password"`
}

type UserFilter struct {
	Name      *string    `json:"name"`
	Email     *string    `json:"email"`
	Birthday  *time.Time `json:"birthday"`
	Password  *string    `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Limit     *int       `json:"limit"`
	Offset    *int       `json:"offset"`
}

func (u UserRepository) GetAllUsers(filter UserFilter) ([]models.User, error) {
	var users []models.User
	var params []string
	var args []interface{}

	query := `
		SELECT user_id, name, email, birthday, created_at, updated_at
		FROM users WHERE deleted_at IS NULL
	`

	if filter.Name != nil {
		args = append(args, *filter.Name)
		params = append(params, fmt.Sprintf("name = $%d", len(args)))
	}

	if filter.Email != nil {
		params = append(params, fmt.Sprintf("email = $%d", len(args)+1))
		args = append(args, *filter.Email)
	}

	if filter.Birthday != nil {
		params = append(params, fmt.Sprintf("birthday = $%d", len(args)+1))
		args = append(args, *filter.Birthday)
	}

	if filter.CreatedAt != nil {
		params = append(params, fmt.Sprintf("created_at = $%d", len(args)+1))
		args = append(args, *filter.CreatedAt)
	}

	if filter.UpdatedAt != nil {
		params = append(params, fmt.Sprintf("updated_at = $%d", len(args)+1))
		args = append(args, *filter.UpdatedAt)
	}

	if len(params) > 0 {
		query += strings.Join(params, " AND ")
	}

	if filter.Limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *filter.Limit)
	}

	if filter.Offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *filter.Offset)
	}

	// fmt.Println("Executing query:", query)
	// fmt.Println("With arguments:", args)

	rows, err := u.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed while querying: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed while scanning data to slice: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return users, nil
}

func (u UserRepository) CreateUser(user *models.User) (*string, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO users (user_id, name, email, birthday, password)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := u.db.Exec(query, id, user.Name, user.Email, user.Birthday, user.Password)
	if err != nil {
		return nil, fmt.Errorf("creating user failed: %v", err)
	}

	return &id, nil
}

func (u UserRepository) GetUser(userID string) (*models.User, error) {
	query := `
		SELECT user_id, name, email, birthday, password, created_at, updated_at
		FROM users
		WHERE deleted_at IS NULL AND user_id = $1
	`
	row := u.db.QueryRow(query, userID)
	user := models.User{}

	err := row.Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}
		return nil, fmt.Errorf("getting user failed: %v", err)
	}

	return &user, nil
}

func (u UserRepository) DeleteUser(userId string) error {
	query := `
		UPDATE users 
		SET deleted_at = CURRENT_TIMESTAMP 
		WHERE user_id = $1 AND deleted_at IS NULL
	`

	if _, err := u.db.Exec(query, userId); err != nil {
		if sql.ErrNoRows == err {
			return errors.New("not found")
		}
		return fmt.Errorf("failed to delete user by this id: %v", err)
	}

	return nil
}

func (u UserRepository) UpdateUser(userId string, updateFilter UpdateUser) error {
	var params []string
	var args []interface{}

	query := `
		SELECT user_id
		FROM users
		WHERE deleted_at IS NULL AND user_id = $1
	`

	if err := u.db.QueryRow(query, userId).Err(); err != nil {
		return fmt.Errorf("user by this id not found: %v", err)
	}

	query = `
		UPDATE users SET updated_at = CURRENT_TIMESTAMP,
	`

	if updateFilter.Name != nil {
		params = append(params, fmt.Sprintf("name = $%d", len(args)+1))
		args = append(args, *updateFilter.Name)
	}

	if updateFilter.Email != nil {
		params = append(params, fmt.Sprintf("email = $%d", len(args)+1))
		args = append(args, *updateFilter.Email)
	}

	if updateFilter.Birthday != nil {
		params = append(params, fmt.Sprintf("birthday = $%d", len(args)+1))
		args = append(args, *updateFilter.Birthday)
	}

	if updateFilter.Password != nil {
		params = append(params, fmt.Sprintf("password = $%d", len(args)+1))
		args = append(args, *updateFilter.Password)
	}

	if len(params) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, userId)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE user_id = $%d AND deleted_at IS NULL", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)

	if _, err := u.db.Exec(query, args...); err != nil {
		return fmt.Errorf("failed executing query: %v", err)
	}
	return nil
}
