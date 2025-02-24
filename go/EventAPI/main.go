package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/utkarsh-singh1/project/go/EventAPI/models"
)

func main() {

	// gin.Default returns an Engine instance
	server := gin.Default()

	// Starts the webserver
	server.Run(":8080")

	// Creating the Routers and the Handler function
	server.GET("/events", getAllEvent)

	server.POST("/events", createNewEvent)
}


func getAllEvent(c *gin.Context) {

	events := models.GetAllEvent()
	c.JSON(http.StatusOK, events)

}

func createNewEvent(c *gin.Context) {

	var event models.Event

	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest,fmt.Sprintln("Wrong Request Sent By User"))
		return
	}
	
	event.ID = 1

	c.JSON(http.StatusCreated, fmt.Sprintf("The current event is created and your registered event info is %v",event))
}
