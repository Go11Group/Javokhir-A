package handlers

import (
	"net/http"

	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/dtos"
	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/services"
	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	CourseService *services.CourseService
}

func NewCourseHandler(courseService *services.CourseService) *CourseHandler {
	return &CourseHandler{
		CourseService: courseService,
	}
}

func (ch *CourseHandler) CreateCourse(c *gin.Context) {
	var courseData dtos.CourseCreateDTO

	if err := c.BindJSON(&courseData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding JSON data failed: " + err.Error()})
		return
	}

	course, err := ch.CourseService.CreateCourse(courseData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"course": course})
}

func (ch *CourseHandler) GetCourse(c *gin.Context) {
	courseID := c.Param("id")

	course, err := ch.CourseService.GetCourse(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"course": course})
}

func (ch *CourseHandler) UpdateCourse(c *gin.Context) {
	courseID := c.Param("id")

	var courseData dtos.CourseCreateDTO
	if err := c.BindJSON(&courseData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding JSON data failed: " + err.Error()})
		return
	}

	err := ch.CourseService.UpdateCourse(courseID, courseData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

func (ch *CourseHandler) DeleteCourse(c *gin.Context) {
	courseID := c.Param("id")

	err := ch.CourseService.DeleteCourse(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
