package handlers

import (
	"net/http"
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/repositories"
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CourseHandlerPlan interface {
	CreateCourse(c *gin.Context)
	GetCourseByID(c *gin.Context)
	UpdateCourse(c *gin.Context)
	DeleteCourse(c *gin.Context)
	GetAllCourses(c *gin.Context)
}

type CourseHandler struct {
	CourseService *services.CourseService
}

func NewCourseHandler(service *services.CourseService) *CourseHandler {
	return &CourseHandler{
		CourseService: service,
	}
}

func (ch CourseHandler) CreateCourse(c *gin.Context) {
	var cc services.CourseCreate
	var errCode = http.StatusBadRequest

	if err := c.BindJSON(&cc); err != nil {
		c.JSON(errCode, gin.H{"error": "failed binding json to course struct: " + err.Error()})
		return
	}

	res, err := ch.CourseService.CreateCourse(&cc)
	if err != nil {
		c.JSON(errCode, gin.H{"error": "creating course failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &res)
}

func (ch CourseHandler) GetCourseByID(c *gin.Context) {
	id := c.Param("id")

	cres, err := ch.CourseService.GetCourseByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "getting response from course service failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, cres)
}

func (ch CourseHandler) UpdateCourse(c *gin.Context) {
	updateCourse := services.UpdateCourse{}
	id := c.Param("id")

	if err := c.BindJSON(&updateCourse); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "binding to json failed: " + err.Error()})
		return
	}

	cId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not valid: " + err.Error()})
		return
	}

	res, err := ch.CourseService.UpdateCourse(updateCourse, cId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "updated course failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ch CourseHandler) DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	errCode := http.StatusBadRequest

	if err := ch.CourseService.DeleteCourse(id); err != nil {
		if err.Error() == "not found" {
			errCode = http.StatusNotFound
		}

		c.JSON(errCode, gin.H{"error": "deleting course failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "course successfully deleted"})
}

func (ch CourseHandler) GetAllCourses(c *gin.Context) {
	filter := repositories.CourseFilter{}

	if err := c.BindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding to json failed: " + err.Error()})
		return
	}

	courses, err := ch.CourseService.GetAllCourses(&filter)
	if err != nil {
		if err.Error() == "timeout" {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": "getting courses time exceeded, use limit to limit the data"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "getting all courses failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusFound, &courses)
}

func (ch CourseHandler) GetCourseByUser(c *gin.Context) {
	id := c.Param("id")
	uuID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is not valid"})
		return
	}

	userCourses, err := ch.CourseService.GetCourseByUser(uuID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get course: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, userCourses)
}

func (ch CourseHandler) GetEnroleldUsersByCourse(c *gin.Context) {
	id := c.Param("id")
	courseID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course id"})
		return
	}

	enUsers, err := ch.CourseService.GetEnroleldUsersByCourse(courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to fetch data or no records found:" + err.Error()})
		return
	}

	c.JSON(http.StatusFound, enUsers)
}

func (ch *CourseHandler) GetMostPopularCoursesHandler(c *gin.Context) {
	var timePeriod repositories.TimePeriod

	if err := c.ShouldBindJSON(&timePeriod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON input"})
		return
	}

	startDate, err := time.Parse("2006-01-02", timePeriod.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format"})
		return
	}

	endDate, err := time.Parse("2006-01-02", timePeriod.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format"})
		return
	}

	response, err := ch.CourseService.GetMostPopularCourses(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, response)
}
