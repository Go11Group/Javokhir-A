package app

import (
	"github.com/Go11Group/Javokhir-A/homework/lesson36/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupHandlers(router *gin.Engine, userSer *services.UsersService, problemSer *services.ProblemService) {
	user := router.Group("/user")

	user.GET("/:id", userSer.GetUser)
	user.PUT("/:id", userSer.UpdateUser)
	user.DELETE("/:id", userSer.DeleteUser)
	user.GET("/getall", userSer.GetAllUsers)

	problem := router.Group("/problem")

	problem.GET("/:id", problemSer.GetProblem)
	problem.PUT("/:id", problemSer.UpdateProblem)
	problem.DELETE("/:id", problemSer.DeleteProblem)
	problem.GET("s/", problemSer.GetAllProblems)

	// solved_problem := router.Group("/solved_problem")

	// solved_problem.GET("/:")
	// router.POST("/user", userSer.CreateUser)
	// router.PUT("/user/:id", userSer.UpdateUser)
	// router.DELETE("/user/:id", userSer.DeleteUser)

	// router.GET("/problems", problemSer.GetAllProblems)
	// router.GET("/problem/:id", problemSer.GetProblem)
	// router.POST("/problem", problemSer.CreateProblem)
	// router.DELETE("/problem/:id", problemSer.DeleteProblem)
}
