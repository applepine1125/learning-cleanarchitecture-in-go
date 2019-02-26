package infra

import (
	"github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/interfaces/controller"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controller.NewUserController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.GetUser(c) })
	router.GET("/users", func(c *gin.Context) { userController.GetUsers(c) })

	Router = router
}
