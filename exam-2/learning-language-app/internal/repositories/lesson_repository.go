package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
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

func NewLessonRepository(db *sql.DB) *LessonRepository {
	return &LessonRepository{
		db: db,
	}
}

func (l *LessonRepository) CreateLesson(lesson *CreateLesson, courseID uuid.UUID) error {
	lessonId := uuid.New()
	query := `
		INSERT INTO lessons(lesson_id, course_id, title, content)
		VALUES($1, $2, $3, $4)
	`
	_, err := l.db.Exec(query, lessonId, courseID, lesson.Title, lesson.Content)

	return err
}

func (l *LessonRepository) GetLessonByID(lessonID uuid.UUID) (*models.Lesson, error) {
	qury := `
		SELECT lesson_id, course_id, title, content, created_at, updated_at
		FROM lessons WHERE deleted_at IS NULL AND lesson_id = $1
	`

	row := l.db.QueryRow(qury, lessonID)
	var lesson models.Lesson

	err := row.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content)

	if err != nil {
		return nil, errors.New("NotFound")
	}

	return &lesson, nil
}

func (l *LessonRepository) UpdateLesson(lf UpdateLesson, lessonID uuid.UUID) error {
	query := `
		UPDATE lessons SET 
	`
	conditions := []string{}
	args := []interface{}{}

	if lf.Title != nil {
		conditions = append(conditions, fmt.Sprintf("title = $%d", len(args)+1))
		args = append(args, *lf.Title)
	}

	if lf.Content != nil {
		conditions = append(conditions, fmt.Sprintf("content = $%d", len(args)+1))
		args = append(args, *lf.Content)
	}

	if len(conditions) > 0 {
		query += strings.Join(conditions, " , ") + `, updated_at = CURRENT_TIMESTAMP 
		WHERE deleted_at IS NULL AND lesson_id = $` + strconv.Itoa(len(args)+1)
	} else {
		return errors.New("NoUpdate")
	}
	args = append(args, lessonID)

	_, err := l.db.Exec(query, args...)

	return err
}

func (l *LessonRepository) DeleteLesson(lessonID uuid.UUID) error {
	query := `UPDATE lessons SET deleted_at = CURRENT_TIMESTAMP WHERE deleted_at IS NULL and lesson_id = $1`

	_, err := l.db.Exec(query, lessonID)

	return err
}
func (l *LessonRepository) GetAllLessons(ctx context.Context, lf LessonFilter) ([]*models.Lesson, error) {
	var conditions []string
	var args []interface{}

	query := `
        SELECT lesson_id, course_id, title, content, created_at, updated_at FROM lessons
        WHERE deleted_at IS NULL   
    `

	if lf.CourseID != nil {
		conditions = append(conditions, fmt.Sprintf("course_id = $%d", len(args)+1))
		args = append(args, *lf.CourseID)
	}
	if lf.Title != nil {
		conditions = append(conditions, fmt.Sprintf("title = $%d", len(args)+1))
		args = append(args, *lf.Title)
	}
	if lf.Content != nil {
		conditions = append(conditions, fmt.Sprintf("content = $%d", len(args)+1))
		args = append(args, *lf.Content)
	}
	if lf.CreatedAt != nil {
		conditions = append(conditions, "DATE(created_at) = $"+strconv.Itoa(len(args)+1))
		args = append(args, lf.CreatedAt.Format("2006-01-02"))
	}
	if lf.UpdatedAt != nil {
		conditions = append(conditions, "DATE(updated_at) = $"+strconv.Itoa(len(args)+1))
		args = append(args, lf.UpdatedAt.Format("2006-01-02"))
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	if lf.Limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *lf.Limit)
	}
	if lf.Offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *lf.Offset)
	}

	rows, err := l.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch lessons: %w", err)
	}
	defer rows.Close()

	var lessons []*models.Lesson
	for rows.Next() {
		var lesson models.Lesson
		if err := rows.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt); err != nil {
			return nil, err
		}
		lessons = append(lessons, &lesson)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating over lessons: %w", err)
	}

	return lessons, nil
}

func (l *LessonRepository) GetLessonByCourse(courseID uuid.UUID) (*CourseLessons, error) {
	query := `
		SELECT lesson_id, title, content 
		FROM lessons
		WHERE course_id = $1 AND deleted_at IS NULL
	`

	rows, err := l.db.Query(query, courseID)
	if err != nil {
		return nil, fmt.Errorf("no rows found" + err.Error())
	}

	var lessons []Lesson
	for rows.Next() {
		var lesson Lesson
		if err := rows.Scan(&lesson.Lesson_id, &lesson.Title, &lesson.Content); err != nil {
			return nil, fmt.Errorf("faield while iterating through rows")
		}
		lessons = append(lessons, lesson)
	}

	return &CourseLessons{
		CourseID: courseID,
		Lessons:  lessons,
	}, nil

}
