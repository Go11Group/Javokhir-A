package repositories

import (
	"database/sql"
	"fmt"

	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/models"
	"github.com/google/uuid"
)

type LessonRepository struct {
	db *sql.DB
}

// UpdateLesson defines the fields that can be updated in a lesson.
type UpdateLesson struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

func NewLessonRepo(db *sql.DB) *LessonRepository {
	return &LessonRepository{
		db: db,
	}
}

func (lr *LessonRepository) CreateLesson(lesson *models.Lesson) (*string, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO lessons (lesson_id, course_id, title, content)
		VALUES ($1, $2, $3, $4)
	`
	_, err := lr.db.Exec(query, id, lesson.CourseID, lesson.Title, lesson.Content)
	if err != nil {
		return nil, fmt.Errorf("creating lesson failed: %v", err)
	}

	return &id, nil
}

func (lr *LessonRepository) GetLesson(lessonID string) (*models.Lesson, error) {
	query := `
		SELECT lesson_id, course_id, title, content, created_at, updated_at, deleted_at
		FROM lessons
		WHERE deleted_at IS NULL AND lesson_id = $1
	`
	row := lr.db.QueryRow(query, lessonID)
	lesson := models.Lesson{}

	err := row.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
	if err != nil {
		return nil, fmt.Errorf("getting lesson failed: %v", err)
	}

	return &lesson, nil
}

func (lr *LessonRepository) UpdateLesson(lessonID string, lesson *models.Lesson) error {
	query := `
		UPDATE lessons 
		SET title = $1, content = $2, updated_at = CURRENT_TIMESTAMP
		WHERE lesson_id = $3 AND deleted_at IS NULL
	`
	_, err := lr.db.Exec(query, lesson.Title, lesson.Content, lessonID)
	if err != nil {
		return fmt.Errorf("updating lesson failed: %v", err)
	}

	return nil
}

func (lr *LessonRepository) DeleteLesson(lessonID string) error {
	query := `
		UPDATE lessons 
		SET deleted_at = CURRENT_TIMESTAMP 
		WHERE lesson_id = $1 AND deleted_at IS NULL
	`
	_, err := lr.db.Exec(query, lessonID)
	if err != nil {
		return fmt.Errorf("deleting lesson failed: %v", err)
	}

	return nil
}
