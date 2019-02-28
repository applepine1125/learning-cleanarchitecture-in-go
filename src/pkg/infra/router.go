package infra

import (
	"github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/interfaces/controller"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controller.NewUserController(NewSqlHandler())

	taskController := controller.NewTaskController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.GetUser(c) })
	router.GET("/users", func(c *gin.Context) { userController.GetUsers(c) })

	router.POST("/tasks", func(c *gin.Context) { taskController.Create(c) })
	router.GET("/tasks/:id", func(c *gin.Context) { taskController.GetTask(c) })
	router.GET("/tasks", func(c *gin.Context) { taskController.GetTasks(c) })
	Router = router
}
