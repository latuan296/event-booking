package main

import (
	"exmpale.com/event-booking/db"
	"exmpale.com/event-booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoute(server)

	server.Run(":8080") // localhost:8080

}
