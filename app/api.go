package app

import (
	"go_taskapi/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

//create a new Task

func NewTaskHandler(c *gin.Context){
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask);err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	newTask.ID = xid.New().String()
	model.Tasks = append(model.Tasks,newTask)
	c.JSON(http.StatusOK, newTask)
	


}

//retrieving all the list of tasks

func GetTaskHandler(c *gin.Context){
	c.JSON(http.StatusOK,model.Tasks)
}


//updating task

func UpdateTaskHandler(c *gin.Context){
	id:= c.Param("id")
	var task model.Task

	if err:= c.ShouldBindJSON(&task); err !=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	index := -1
	for i:=0;i < len(model.Tasks);i++{
		if model.Tasks[i].ID == id{
			index =i
		}
	}
	if index == -1{
		c.JSON(http.StatusNotFound,gin.H{"error":"Task not found"})
		return
	}
	task.ID = xid.New().String()
	model.Tasks[index] = task
	c.JSON(http.StatusOK,task)
}

//deleting

func DeleteTaskHandler(c *gin.Context){
	id := c.Param("id")

	index :=-1

	for i:=0; i<len(model.Tasks);i++{
		if model.Tasks[i].ID == id{
			index =i

		}
		
	}
	if index == -1{
		c.JSON(http.StatusNotFound,gin.H{"error": "Task not found"})
	}
	model.Tasks = append(model.Tasks[:index],model.Tasks[index+1:]...)
	c.JSON(http.StatusOK,gin.H{"message":" Task has been deleted"})
}