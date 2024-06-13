package dtos

import "time"

type LessonCreateDTO struct {
	CourseID string `json:"course_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type LessonResponseDTO struct {
	LessonID  string    `json:"lesson_id"`
	CourseID  string    `json:"course_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type LessonUpdateDTO struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	UpdatedAt string `json:"updated_at"`
}
