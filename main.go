package main

import (
	"github.com/RafaelRochaS/journey-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	server.SetTrustedProxies([]string{"192.168.15.1"})

	setMiddleware(server)
	routes.InitializeRoutes(server)

	server.Run(":8080")
}

func setMiddleware(server *gin.Engine) {
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
}
