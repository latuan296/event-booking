package routes

import (
	"exmpale.com/event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// Middleware authentication
	authenticatedRoutes := server.Group("/")
	authenticatedRoutes.Use(middlewares.Authenticate)
	authenticatedRoutes.POST("/events", createEvent)
	authenticatedRoutes.PUT("/events/:id", updateEvent)
	authenticatedRoutes.DELETE("/events/:id", deleteEvent)
	authenticatedRoutes.POST("/events/:id/register", registerForEvent)
	authenticatedRoutes.DELETE("/events/:id/register", cancelRegistration)

	// USER ROUTES
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
