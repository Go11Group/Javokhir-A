package models

import "time"

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Problem struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Complexity  string    `json:"compexity"`
	Category    Category  `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_ate"`
	DeletedAt   time.Time `json:"deleted_at"`
}
