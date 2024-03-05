package main

import (
	"go_taskapi/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Task Management Api"})

}

func main() {
	router := gin.Default()

	router.GET("/", HomeHandler)
	router.GET("/tasks", app.GetTaskHandler)
	router.POST("task", app.NewTaskHandler)
	router.PUT("/task/:id", app.UpdateTaskHandler)
	router.DELETE("/task/:id", app.DeleteTaskHandler)
	err := router.Run()
	if err != nil {
		return
	}

}
