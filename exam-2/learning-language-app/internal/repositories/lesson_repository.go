package repositories

import (
	"database/sql"
	"errors"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/models"
)

// Plan for lesson repo
type LessonRepositoryPlan interface {
	CreateLesson(lesson models.Lesson) error
	GetLessonByID(lessonID string) (models.Lesson, error)
	UpdateLesson(lesson models.Lesson) error
	DeleteLesson(lessonID string) error
	GetAllLessons() ([]models.Lesson, error)
}

type LessonRepository struct {
	db *sql.DB
}

func NewLessonRepository(db *sql.DB) *LessonRepository {
	return &LessonRepository{
		db: db,
	}
}

func (r *LessonRepository) CreateLesson(lesson models.Lesson) error {
	query := `INSERT INTO lessons (lesson_id, course_id, title, content)
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, lesson.LessonID, lesson.CourseID, lesson.Title, lesson.Content)
	return err
}

func (r *LessonRepository) GetLessonByID(lessonID string) (models.Lesson, error) {
	query := `SELECT lesson_id, course_id, title, content, created_at, updated_at
              FROM lessons WHERE lesson_id = $1`
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
	query := `UPDATE lessons SET course_id = $1, title = $2, content = $3, updated_at = $4 WHERE lesson_id = $5`
	_, err := r.db.Exec(query, lesson.CourseID, lesson.Title, lesson.Content, lesson.UpdatedAt, lesson.LessonID)
	return err
}

func (r *LessonRepository) DeleteLesson(lessonID string) error {
	query := `DELETE FROM lessons WHERE lesson_id = $1`
	_, err := r.db.Exec(query, lessonID)
	return err
}

func (r *LessonRepository) GetAllLessons() ([]models.Lesson, error) {
	query := `SELECT lesson_id, course_id, title, content, created_at, updated_at, deleted_at FROM lessons`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []models.Lesson
	for rows.Next() {
		var lesson models.Lesson
		err := rows.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}
