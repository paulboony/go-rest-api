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
	router := SetupRoutes()
	router.Run(":8080")
}

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	health.Route(router, healthController)
	task.Route(router, taskController)
	return router
}
