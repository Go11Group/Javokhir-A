package dtos

import "time"

type CourseCreateDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CourseUpdateDTO struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
type CourseResponseDTO struct {
	CourseID    string    `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
