package repositories

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Go11Group/Javokhir-A/homework/lesson36/internal/models"
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
	Name     *string  `json:"name"`
	UserName *string  `json:"user_name"`
	Rank     *float64 `json:"rank"`
	Password *string  `json:"password"`
}

type UserFilter struct {
	Name      *string    `json:"name"`
	UserName  *string    `json:"user_name"`
	Rank      *float64   `json:"rank"`
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
		SELECT id, name, user_name, rank, created_at, updated_at
		FROM users WHERE deleted_at IS NULL
	`

	if filter.Name != nil {
		params = append(params, fmt.Sprintf("name = $%d", len(args)+1))
		args = append(args, *filter.Name)
	}

	if filter.UserName != nil {
		params = append(params, fmt.Sprintf("user_name = $%d", len(args)+1))
		args = append(args, *filter.UserName)
	}

	if filter.Rank != nil {
		params = append(params, fmt.Sprintf("rank = $%d", len(args)+1))
		args = append(args, *filter.Rank)
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
		if err := rows.Scan(&user.Id, &user.Name, &user.UserName, &user.Rank, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed while scanning data to slice: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return users, nil
}

// func (u UserRepository) GetAllUsers(filter UserFilter) error{
// 	var users []models.User

// 	param := make(map[string]interface{})
// 	args := []interface

// 	query := `
// 		SELECT
// 			id,
// 			Name,
// 			UserName,
// 			Password,
// 			Rank,
// 			CreatedAt,
// 			UpdatedAt
// 		FROM users
// 		WEHRE deleted_at IS NULL
// 	`
// 	f := ""

// 	if filter.Name != nil{
// 		param["name"] = filter.Name
// 		f += "AND name = $"
// 	}
// 	if filter.Rank != nil{
// 		param["rank"] += filter.Rank
// 		f += "AND name = $"

// 	}

// 	if filter.UserName != nil{
// 		param["user_name"] = filter.UserName
// 		f += "AND user_name = $"

// 	}

// 	if filter.CreatedAt != nil{
// 		param["created_at"] = filter.CreatedAt
// 		f += "AND created_at = $"

// 	}

// 	if filter.UpdatedAt != nil{
// 		param["upadted_at"] = filter.UpdatedAt
// 		f += "AND updated_at = $"
// 	}

// 	if filter.Limit != nil{
// 		param["limit"] += filter.Limit
// 		f += "LIMIT $"
// 	}

// 	if filter.Offset != nil{
// 		param["offset"] += filter.Offset
// 		f += "OFFSET $"
// 	}
// 	i := 1
// 	for

// 	u.db.Exec(query, )

// }

func (u UserRepository) CreateUser(user *models.User) error {
	id := uuid.NewString()
	query := `
		INSERT INTO users (id, name, user_name, password)
		VALUES ($1, $2, $3, $4)
	`
	_, err := u.db.Exec(query, id, user.Name, user.UserName, user.Password)
	if err != nil {
		return fmt.Errorf("creating user failed: %v", err)
	}

	return nil
}

func (u UserRepository) GetUser(userID string) (*models.User, error) {
	query := `
		SELECT id, name, user_name, password, rank 
		FROM users
		WHERE deleted_at IS NULL AND id = $1
	`
	row := u.db.QueryRow(query, userID)
	user := models.User{}

	err := row.Scan(&user.Id, &user.Name, &user.UserName, &user.Password, &user.Rank)
	if err != nil {
		return nil, fmt.Errorf("getting user failed: %v", err)
	}

	return &user, nil
}

func (u UserRepository) DeleteUser(userId string) error {
	query := `
		UPDATE users 
		SET deleted_at = CURRENT_TIMESTAMP 
		WHERE id = $1 AND deleted_at IS NULL
	`

	if _, err := u.db.Exec(query, userId); err != nil {
		return fmt.Errorf("failed to delete user by this id: %v", err)
	}

	return nil
}

func (u UserRepository) UpdateUser(userId string, updateFilter UpdateUser) error {
	var params []string
	var args []interface{}

	query := `
		SELECT id
		FROM users
		WHERE deleted_at IS NULL AND id = $1
	`

	if err := u.db.QueryRow(query, userId).Err(); err != nil {
		return fmt.Errorf("user by this id not found: %v", err)
	}

	query = `
		UPDATE users SET 
	`

	if updateFilter.Name != nil {
		params = append(params, fmt.Sprintf("name = $%d", len(args)+1))
		args = append(args, *updateFilter.Name)
	}

	if updateFilter.Rank != nil {
		params = append(params, fmt.Sprintf("rank = $%d", len(args)+1))
		args = append(args, *updateFilter.Rank)
	}

	if updateFilter.UserName != nil {
		params = append(params, fmt.Sprintf("user_name = $%d", len(args)+1))
		args = append(args, *updateFilter.UserName)
	}

	if updateFilter.Password != nil {
		params = append(params, fmt.Sprintf("password = $%d", len(args)+1))
		args = append(args, *updateFilter.Password)
	}

	if len(params) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, userId)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE id = $%d AND deleted_at IS NULL", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)

	if _, err := u.db.Exec(query, args...); err != nil {
		return fmt.Errorf("failed executing query: %v", err)
	}
	return nil
}
