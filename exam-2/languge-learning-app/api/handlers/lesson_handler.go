package handlers

import (
	"net/http"

	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/dtos"
	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/services"
	"github.com/gin-gonic/gin"
)

type LessonHandler struct {
	LessonService *services.LessonService
}

func NewLessonHandler(lessonService *services.LessonService) *LessonHandler {
	return &LessonHandler{
		LessonService: lessonService,
	}
}

func (lh *LessonHandler) CreateLesson(c *gin.Context) {
	var lessonData dtos.LessonCreateDTO

	if err := c.BindJSON(&lessonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding JSON data failed: " + err.Error()})
		return
	}

	lesson, err := lh.LessonService.CreateLesson(lessonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"lesson": lesson})
}

func (lh *LessonHandler) GetLesson(c *gin.Context) {
	lessonID := c.Param("id")

	lesson, err := lh.LessonService.GetLesson(lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"lesson": lesson})
}

func (lh *LessonHandler) UpdateLesson(c *gin.Context) {
	lessonID := c.Param("id")

	var lessonData dtos.LessonUpdateDTO
	if err := c.BindJSON(&lessonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding JSON data failed: " + err.Error()})
		return
	}

	_, err := lh.LessonService.UpdateLesson(lessonID, lessonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lesson updated successfully"})
}

func (lh *LessonHandler) DeleteLesson(c *gin.Context) {
	lessonID := c.Param("id")

	err := lh.LessonService.DeleteLesson(lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lesson deleted successfully"})
}
