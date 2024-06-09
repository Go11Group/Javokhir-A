package models

import (
	"time"
)

type User struct {
	Id        string
	Name      string    `json:"name"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	Rank      float64   `json:"rank"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_ate"`
	DeletedAt time.Time `json:"deleted_at"`
}
