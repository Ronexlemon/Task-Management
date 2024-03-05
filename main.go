package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"message":"Task Management Api"})

}

func main(){
	router := gin.Default()

	router.GET("/",HomeHandler)
	err:= router.Run()
	if err !=nil{
		return
	}


}