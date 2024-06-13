package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/models"
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
	CourseID  *string    `json:"course_id"`
	Title     *string    `json:"title"`
	Content   *string    `json:"content"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Limit     *int       `json:"limit"`
	Offset    *int       `json:"offset"`
}

type LessonRepository struct {
	db *sql.DB
}

func NewLessonRepository(db *sql.DB) *LessonRepository {
	return &LessonRepository{
		db: db,
	}
}

func (r *LessonRepository) CreateLesson(lesson *models.Lesson) error {
	query := `INSERT INTO lessons (lesson_id, course_id, title, content, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, lesson.LessonID, lesson.CourseID, lesson.Title, lesson.Content, lesson.CreatedAt, lesson.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to execute the query: %w", err)
	}
	return nil
}

func (r *LessonRepository) GetLessonByID(lessonID string) (models.Lesson, error) {
	query := `SELECT lesson_id, course_id, title, content, created_at, updated_at
              FROM lessons WHERE lesson_id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRow(query, lessonID)

	var lesson models.Lesson
	err := row.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return lesson, errors.New("lesson not found")
		}
		return lesson, err
	}

	return lesson, nil
}

func (r *LessonRepository) UpdateLesson(lesson models.Lesson) error {
	query := `UPDATE lessons SET course_id = $1, title = $2, content = $3, updated_at = $4 WHERE lesson_id = $5 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, lesson.CourseID, lesson.Title, lesson.Content, time.Now(), lesson.LessonID)
	return err
}

func (r *LessonRepository) DeleteLesson(lessonID string) error {
	query := `UPDATE lessons SET deleted_at = CURRENT_TIMESTAMP WHERE lesson_id = $1 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, lessonID)
	return err
}

func (r *LessonRepository) GetAllLessons(f *LessonFilter, ctx context.Context) ([]models.Lesson, error) {
	query := `SELECT lesson_id, course_id, title, content, created_at, updated_at
	          FROM lessons WHERE deleted_at IS NULL`

	var conditions []string
	var args []interface{}

	if f.CourseID != nil {
		conditions = append(conditions, fmt.Sprintf("course_id = $%d", len(args)+1))
		args = append(args, *f.CourseID)
	}
	if f.Title != nil {
		conditions = append(conditions, fmt.Sprintf("title = $%d", len(args)+1))
		args = append(args, *f.Title)
	}
	if f.Content != nil {
		conditions = append(conditions, fmt.Sprintf("content = $%d", len(args)+1))
		args = append(args, *f.Content)
	}
	if f.CreatedAt != nil {
		conditions = append(conditions, fmt.Sprintf("created_at >= $%d AND created_at < $%d", len(args)+1, len(args)+2))
		args = append(args, *f.CreatedAt, f.CreatedAt.AddDate(0, 0, 1))
	}
	if f.UpdatedAt != nil {
		conditions = append(conditions, fmt.Sprintf("updated_at >= $%d AND updated_at < $%d", len(args)+1, len(args)+2))
		args = append(args, *f.UpdatedAt, f.UpdatedAt.AddDate(0, 0, 1))
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	if f.Limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *f.Limit)
	}

	if f.Offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *f.Offset)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []models.Lesson
	for rows.Next() {
		var lesson models.Lesson
		err := rows.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return lessons, nil
}
