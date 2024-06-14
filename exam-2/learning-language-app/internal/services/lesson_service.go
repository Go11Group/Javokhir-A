package services

import (
	"context"
	"errors"
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/models"
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/repositories"
	"github.com/google/uuid"
)

type LessonServicePlan interface {
	CreateLesson(courseID uuid.UUID, title, content string) (*models.Lesson, error)
	GetLessonByID(lessonID uuid.UUID) (*models.Lesson, error)
	UpdateLesson(lessonID uuid.UUID, title, content string) (*models.Lesson, error)
	DeleteLesson(lessonID uuid.UUID) error
	GetAllLessons(filter *repositories.LessonFilter) ([]*models.Lesson, error)
}

type LessonService struct {
	LessonRepository *repositories.LessonRepository
}

func NewLessonService(repo *repositories.LessonRepository) *LessonService {
	return &LessonService{
		LessonRepository: repo,
	}
}

func (ls *LessonService) CreateLesson(courseID uuid.UUID, title, content string) (*models.Lesson, error) {
	lesson := &models.Lesson{
		LessonID: uuid.New(),
		CourseID: courseID,
		Title:    title,
		Content:  content,
	}
	createLesosn := repositories.CreateLesson{
		LessonID: &lesson.LessonID,
		CourseID: &courseID,
		Title:    &title,
		Content:  &content,
	}
	return lesson, ls.LessonRepository.CreateLesson(&createLesosn, courseID)
}

func (ls *LessonService) GetLessonByID(lessonID uuid.UUID) (*models.Lesson, error) {
	return ls.LessonRepository.GetLessonByID(lessonID)
}

func (ls *LessonService) UpdateLesson(lessonID uuid.UUID, title, content string) (*models.Lesson, error) {
	lesson := models.Lesson{
		LessonID:  lessonID,
		Title:     title,
		Content:   content,
		UpdatedAt: time.Now(),
	}
	err := ls.LessonRepository.UpdateLesson(repositories.UpdateLesson{
		Title:   &title,
		Content: &content,
	}, lessonID)

	return &lesson, err
}

func (ls *LessonService) DeleteLesson(lessonID uuid.UUID) error {
	return ls.LessonRepository.DeleteLesson(lessonID)
}

func (ls *LessonService) GetAllLessons(filter *repositories.LessonFilter) ([]*models.Lesson, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	resultCh := make(chan []*models.Lesson)
	errCh := make(chan error)
	defer cancel()

	go func() {
		lessons, err := ls.LessonRepository.GetAllLessons(ctx, *filter)
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- lessons
	}()

	select {
	case <-ctx.Done():
		return nil, errors.New("TimeOut")
	case <-errCh:
		return nil, <-errCh
	case <-resultCh:
		return <-resultCh, nil
	}
}

func (ls *LessonService) GetLessonsByCourse(courseID uuid.UUID) (*repositories.CourseLessons, error) {

	return ls.LessonRepository.GetLessonByCourse(courseID)

}
