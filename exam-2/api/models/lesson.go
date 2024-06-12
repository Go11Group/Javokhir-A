package models

import "time"

type Lesson struct {
	LessonId  string    `json:"lesson_id"`
	CourseId  string    `json:"course_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
