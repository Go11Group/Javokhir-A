package models

type User struct {
	Id                string `json:"id"`
	UserName          string `json:"user_name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	PhoneNumber       string `json:"phone_number"`
	Address           string `json:"address"`
	ProfilePictureUrl string `json:"profile_picture_url"`
	Roles             string `json:"roles"`
	LastLogin         string `json:"last_login"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	Active            bool   `json:"active"`
}
