package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/models"
	"github.com/google/uuid"
)

// Plan for lesson repo
type LessonRepositoryPlan interface {
	CreateLesson(lesson models.Lesson) error
	GetLessonByID(lessonID string) (models.Lesson, error)
	UpdateLesson(lesson models.Lesson) error
	DeleteLesson(lessonID string) error
	GetAllLessons(filter *LessonFilter, ctx context.Context) ([]models.Lesson, error)
}

type LessonFilter struct {
	CourseID  *uuid.UUID `json:"course_id"`
	Title     *string    `json:"title"`
	Content   *string    `json:"content"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Limit     *int       `json:"limit"`
	Offset    *int       `json:"offset"`
}

type CreateLesson struct {
	LessonID *uuid.UUID `json:"lesson_id"`
	CourseID *uuid.UUID `json:"course_id"`
	Title    *string    `json:"title"`
	Content  *string    `json:"content"`
}

type UpdateLesson struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

type CourseLessons struct {
	CourseID uuid.UUID `json:"course_id"`
	Lessons  []Lesson  `json:"lessons"`
}

type Lesson struct {
	Lesson_id uuid.UUID `json:"lesson_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

type LessonRepository struct {
	db *sql.DB
}
