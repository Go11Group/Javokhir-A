package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Go11Group/Javokhir-A/homework/lesson36/internal/models"
	"github.com/Go11Group/Javokhir-A/homework/lesson36/internal/repositories"
	"github.com/gin-gonic/gin"
)

type ProblemService struct {
	problemRepo *repositories.ProblemRepository
}

func NewProblemService(problemRepo *repositories.ProblemRepository) *ProblemService {
	return &ProblemService{
		problemRepo: problemRepo,
	}
}
func (u *ProblemService) CreateProblem(c *gin.Context) {
	var problem models.Problem
	if err := c.BindJSON(&problem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding request body's json to struct failed: " + err.Error()})
		return
	}

	if err := u.problemRepo.CreateProblem(&problem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Creating problem data in database failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Problem successfully created"})
}

func (u *ProblemService) GetAllProblems(ctx *gin.Context) {
	var filter repositories.ProblemFilter

	if err := ctx.BindJSON(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, &filter)
		log.Println("Getting filter from request failed: " + err.Error())
		return
	}

	problems, err := u.problemRepo.GetAllProblems(filter)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetch problems: " + err.Error()})
	}

	ctx.JSON(http.StatusOK, problems)

}
func (u *ProblemService) GetProblem(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(
			http.StatusBadRequest, gin.H{"error": "check for manual url request"},
		)
		log.Println("id must be string: ")
		return
	}

	problem, err := u.problemRepo.GetProblem(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, problem)

}

func (u *ProblemService) UpdateProblem(ctx *gin.Context) {
	var updatingProblem repositories.UpdateProblem

	if err := ctx.BindJSON(&updatingProblem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "updating problem by query failed: " + err.Error()})
		return
	}

	id := ctx.Param("id")

	if err := u.problemRepo.UpdateProblem(id, updatingProblem); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal updating failed: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user data successfully updated"})

}

func (p *ProblemService) DeleteProblem(c *gin.Context) {

	id := c.Param("id")

	if err := p.problemRepo.DeleteProblem(id); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "deleting problem failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("problem id %s deleted", id)})
}
