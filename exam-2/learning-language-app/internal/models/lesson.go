package models

import (
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	LessonID  uuid.UUID  `json:"lesson_id"`
	CourseID  uuid.UUID  `json:"course_id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
