package routes

import (
	"net/http"

	"exmpale.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not create user"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
