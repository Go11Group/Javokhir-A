package models

import "time"

// keeping some fileds with pointers to indicate these fields are optional or needed to add later on
type User struct {
	UserId       string    `json:"user_id"` // id is required
	FirstName    *string   `json:"first_name"`
	LastName     *string   `json:"last_name"`
	Email        string    `json:"email"` // email is required
	Age          *int      `json:"age"`
	Phone        *string   `json:"phone"`
	Role         string    `json:"role"` // default value for this is "user"
	PasswordHash string    `json:"password_hash"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}
