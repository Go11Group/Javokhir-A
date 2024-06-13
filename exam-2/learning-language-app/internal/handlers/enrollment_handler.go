package handlers

import (
	"net/http"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/models"
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/repositories"
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/services"
	"github.com/gin-gonic/gin"
)

type EnrollmentHandlerPlan interface {
	EnrollUser(c *gin.Context)
	GetEnrollmentByID(c *gin.Context)
	DeleteEnrollment(c *gin.Context)
	GetAllEnrollments(c *gin.Context)
}

type EnrollmentHandler struct {
	EnrollmentService *services.EnrollmentService
}

func NewEnrollmentHandler(service *services.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{
		EnrollmentService: service,
	}
}

func (eh EnrollmentHandler) EnrollUser(c *gin.Context) {
	var enrollment models.Enrollment

	if err := c.BindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding to json failed: " + err.Error()})
		return
	}

	enrollmentID, err := eh.EnrollmentService.EnrollUser(enrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enrolling user failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, enrollmentID)
}

func (eh EnrollmentHandler) GetEnrollmentByID(c *gin.Context) {
	id := c.Param("id")

	enrollment, err := eh.EnrollmentService.GetEnrollmentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "fetching enrollment failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, enrollment)
}

func (eh EnrollmentHandler) DeleteEnrollment(c *gin.Context) {
	id := c.Param("id")

	if err := eh.EnrollmentService.DeleteEnrollment(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "deleting enrollment failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "enrollment successfully deleted"})
}

func (eh EnrollmentHandler) GetAllEnrollments(c *gin.Context) {
	filter := repositories.EnrollmentFilter{}
	if err := c.BindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding to json failed: " + err.Error()})
		return
	}

	enrollments, err := eh.EnrollmentService.GetAllEnrollments(filter)
	if err != nil {
		if err.Error() == "timeout" {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": "getting enrollments time exeeded, use limit to limit the data"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "getting all enrollments failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}
