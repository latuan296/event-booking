package routes

import (
	"net/http"
	"strconv"

	"exmpale.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {

	userID := context.GetInt64("UserID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	err = event.Register(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register the event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registration"})
}

func cancelRegistration(context *gin.Context) {

	userID := context.GetInt64("UserID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	var event models.Event
	event.ID = eventID

	err = event.CancelRegistration(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel the registration"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Cancelled"})
}
