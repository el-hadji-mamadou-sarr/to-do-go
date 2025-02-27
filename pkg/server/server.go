package server

import (
	"to-do-go/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/tasks", handlers.GetTasks)
	r.POST("/tasks", handlers.CreateTask)

	return r
}
