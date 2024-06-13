package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/models"
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/repositories"
	"github.com/google/uuid"
)

type LessonServicePlan interface {
	CreateLesson(lesson *LessonCreate) (*LessonResponse, error)
	GetLessonByID(lessonID string) (*LessonResponse, error)
	UpdateLesson(lesson UpdateLesson, lessonID uuid.UUID) (*LessonResponse, error)
	DeleteLesson(lessonID string) error
	GetAllLessons() ([]models.Lesson, error)
}

type LessonService struct {
	LessonRepository *repositories.LessonRepository
}

type LessonCreate struct {
	CourseID  string    `json:"course_id" binding:"required"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateLesson struct {
	CourseID  string    `json:"course_id" binding:"omitempty"`
	Title     string    `json:"title" binding:"omitempty"`
	Content   string    `json:"content" binding:"omitempty"`
	UpdatedAt time.Time `json:"updated_at" binding:"omitempty"`
}

type LessonResponse struct {
	LessonID  string    `json:"lesson_id"`
	CourseID  string    `json:"course_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewLessonService(repo *repositories.LessonRepository) *LessonService {
	return &LessonService{
		LessonRepository: repo,
	}
}

func (ls LessonService) CreateLesson(lesson *LessonCreate) (*LessonResponse, error) {
	newId := uuid.NewString()
	newLesson := models.Lesson{
		LessonID:  newId,
		CourseID:  lesson.CourseID,
		Title:     lesson.Title,
		Content:   lesson.Content,
		CreatedAt: lesson.CreatedAt,
		UpdatedAt: lesson.UpdatedAt,
	}

	if err := ls.LessonRepository.CreateLesson(&newLesson); err != nil {
		return nil, fmt.Errorf("failed to create lesson: %v", err)
	}

	lessonRes := &LessonResponse{
		LessonID:  newId,
		CourseID:  lesson.CourseID,
		Title:     lesson.Title,
		Content:   lesson.Content,
		CreatedAt: lesson.CreatedAt,
		UpdatedAt: lesson.UpdatedAt,
	}
	return lessonRes, nil
}

func (ls LessonService) GetLessonByID(lessonID string) (*LessonResponse, error) {
	lesson, err := ls.LessonRepository.GetLessonByID(lessonID)
	if err != nil {
		return nil, errors.New("getting lesson failed: " + err.Error())
	}

	res := LessonResponse{
		LessonID:  lesson.LessonID,
		CourseID:  lesson.CourseID,
		Title:     lesson.Title,
		Content:   lesson.Content,
		CreatedAt: lesson.CreatedAt,
		UpdatedAt: lesson.UpdatedAt,
	}

	return &res, nil
}

func (ls LessonService) UpdateLesson(lesson UpdateLesson, lessonID uuid.UUID) (*LessonResponse, error) {
	l, err := ls.LessonRepository.GetLessonByID(lessonID.String())
	if err != nil {
		return nil, errors.New("getting lesson to update failed: " + err.Error())
	}

	lRes := LessonResponse{
		LessonID:  lessonID.String(),
		CourseID:  l.CourseID,
		Title:     l.Title,
		Content:   l.Content,
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
	}

	if !lesson.UpdatedAt.IsZero() {
		l.UpdatedAt = lesson.UpdatedAt
		lRes.UpdatedAt = lesson.UpdatedAt
	}

	if lesson.CourseID != "" {
		l.CourseID = lesson.CourseID
		lRes.CourseID = lesson.CourseID
	}

	if lesson.Title != "" {
		l.Title = lesson.Title
		lRes.Title = lesson.Title
	}

	if lesson.Content != "" {
		l.Content = lesson.Content
		lRes.Content = lesson.Content
	}

	if err := ls.LessonRepository.UpdateLesson(l); err != nil {
		return nil, err
	}

	return &lRes, nil
}

func (ls LessonService) DeleteLesson(lessonID string) error {
	if err := ls.LessonRepository.DeleteLesson(lessonID); err != nil {
		if sql.ErrNoRows == err {
			return fmt.Errorf("not found")
		}
		return fmt.Errorf("deleting lesson failed: " + err.Error())
	}
	return nil
}

func (ls LessonService) GetAllLessons(f *repositories.LessonFilter) ([]models.Lesson, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// Using a channel to handle the result and error
	resultCh := make(chan []models.Lesson)
	errCh := make(chan error)

	go func() {
		lessons, err := ls.LessonRepository.GetAllLessons(f, ctx)
		if err != nil {
			errCh <- fmt.Errorf("getting lessons failed: %v", err)
			return
		}
		resultCh <- lessons
	}()

	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	case lessons := <-resultCh:
		return lessons, nil
	case err := <-errCh:
		return nil, err
	}
}
