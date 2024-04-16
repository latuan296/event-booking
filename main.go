package main

import (
	"fmt"
	"net/http"

	"exmpale.com/event-booking/db"
	"exmpale.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": events})
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event"})
		fmt.Println(err)
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create events"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Created event successfully", "event": event})
}
