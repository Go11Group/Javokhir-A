package dtos

import "time"

// UserCreateDTO. DTO data Transfer Objects. In this API request body should contain specific fileds
// binding is a go tag for gin framework. binding:required tag tells
// gin that incoming data should be inserted.This ensures user input validation while requesting to the server

type UserCreateDTO struct {
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Birthday  string    `json:"birthday" binding:"required"`
	Password  string    `json:"password" binding:"required,min=6"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// omitempty - omits when the value is not inserted to a field and not wanting to get its zero values
type UserUpdateDTO struct {
	Name      *string `json:"name,omitempty"`
	Email     *string `json:"email,omitempty"`
	Birthday  *string `json:"birthday,omitempty"`
	Password  *string `json:"password,omitempty"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type UserResponseDTO struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Birthday  string    `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
