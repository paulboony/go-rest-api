package main

import (
	"github.com/gin-gonic/gin"

	"paulboony/go-rest-api/health"
	"paulboony/go-rest-api/task"
)

var (
	healthController health.HealthController = health.NewController()
	taskService      task.TaskService        = task.NewService()
	taskController   task.TaskController     = task.NewController(taskService)
)

func main() {
	server := gin.Default()
	server.GET("/", healthController.Health)
	server.GET("/tasks", taskController.FindAll)
	server.GET("/tasks/:id", taskController.FindById)
	server.POST("/tasks", taskController.Create)
	server.PATCH("/tasks/:id", taskController.Update)
	server.DELETE("/tasks/:id", taskController.Delete)

	server.Run(":8080")
}
