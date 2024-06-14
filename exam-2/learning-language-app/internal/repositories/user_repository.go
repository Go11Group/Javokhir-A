package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/models"
)

func (u *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (user_id, name, email, birthday, password)
              VALUES ($1, $2, $3, $4, $5)`
	_, err := u.db.Exec(query, user.UserID, user.Name, user.Email, user.Birthday, user.Password)
	if err != nil {
		return fmt.Errorf("failed execute the query")
	}
	return nil
}

func (u *UserRepository) GetUserByID(userID string) (models.User, error) {
	query := `SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at 
              FROM users WHERE user_id = $1 and deleted_at IS NULL`
	row := u.db.QueryRow(query, userID)

	var user models.User
	err := row.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}

func (u *UserRepository) UpdateUser(user models.User) error {
	query := `UPDATE users SET name = $1, email = $2, birthday = $3, password = $4, updated_at = $5 WHERE user_id = $6 and deleted_at IS NULL`
	_, err := u.db.Exec(query, user.Name, user.Email, user.Birthday, user.Password, time.Now(), user.UserID)
	return err
}

func (u *UserRepository) DeleteUser(userID string) error {
	query := `UPDATE FROM users SET deleted_at = CURRENT_TIMESTAMP WHERE user_id = $1 and deleted_at IS NULL`
	_, err := u.db.Exec(query, userID)
	return err
}
func (u *UserRepository) GetAllUsers(f *UserFilter, ctx context.Context) ([]models.User, error) {
	query := `SELECT user_id, name, email, birthday, password, created_at, updated_at FROM users WHERE deleted_at IS NULL`

	var conditions []string
	var args []interface{}

	if f.Name != nil {
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(args)+1))
		args = append(args, *f.Name)
	}
	if f.Email != nil {
		conditions = append(conditions, fmt.Sprintf("email = $%d", len(args)+1))
		args = append(args, *f.Email)
	}
	if f.Birthday != nil {
		conditions = append(conditions, fmt.Sprintf("birthday = $%d", len(args)+1))
		args = append(args, *f.Birthday)
	}
	if f.CreatedAt != nil {
		conditions = append(conditions, fmt.Sprintf("created_at = $%d", len(args)+1))
		args = append(args, *f.CreatedAt)
	}
	if f.UpdatedAt != nil {
		conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)+1))
		args = append(args, *f.UpdatedAt)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	if f.Limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *f.Limit)
	}

	if f.Offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *f.Offset)
	}

	fmt.Println(query, args)
	rows, err := u.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	// fmt.Println(users)
	return users, nil
}
