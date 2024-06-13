package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/models"
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/repositories"
	"github.com/google/uuid"
)

type EnrollmentServicePlan interface {
	EnrollUser(enrollment models.Enrollment) (*models.Enrollment, error)
	GetEnrollmentByID(enrollmentID string) (*models.Enrollment, error)
	DeleteEnrollment(enrollmentID string) error
	GetAllEnrollments(filter *repositories.EnrollmentFilter) ([]models.Enrollment, error)
}

type EnrollmentService struct {
	EnrollmentRepository *repositories.EnrollmentRepository
}

func NewEnrollmentService(repo *repositories.EnrollmentRepository) *EnrollmentService {
	return &EnrollmentService{
		EnrollmentRepository: repo,
	}
}

func (es EnrollmentService) EnrollUser(enrollment models.Enrollment) (*models.Enrollment, error) {
	enrollment.EnrollmentID = uuid.NewString()

	if err := es.EnrollmentRepository.EnrollUser(enrollment); err != nil {
		return nil, err
	}

	return &enrollment, nil
}

func (es EnrollmentService) GetEnrollmentByID(enrollmentID string) (*models.Enrollment, error) {
	enrollment, err := es.EnrollmentRepository.GetEnrollmentByID(enrollmentID)
	if err != nil {
		return nil, errors.New("fetching enrollment failed: " + err.Error())
	}
	return &enrollment, nil
}

func (es EnrollmentService) DeleteEnrollment(enrollmentID string) error {
	if err := es.EnrollmentRepository.DeleteEnrollment(enrollmentID); err != nil {
		return errors.New("deleting enrollment failed: " + err.Error())
	}
	return nil
}

func (es EnrollmentService) GetAllEnrollments(filter repositories.EnrollmentFilter) ([]models.Enrollment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	resultCh := make(chan []models.Enrollment)
	errCh := make(chan error)

	go func() {
		enrollments, err := es.EnrollmentRepository.GetAllEnrollments(&ctx, filter)
		if err != nil {
			errCh <- fmt.Errorf("getting enrollments failed: %v", err)
			return
		}
		resultCh <- enrollments
	}()

	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	case enrollments := <-resultCh:
		return enrollments, nil
	case err := <-errCh:
		return nil, err
	}
}
