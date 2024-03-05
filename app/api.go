package app

import (
	"go_taskapi/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

//create a new Task

func newTaskHandler(c *gin.Context){
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask);err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	newTask.ID = xid.New().String()
	model.Tasks = append(model.Tasks,newTask)
	c.JSON(http.StatusOK, newTask)
	


}