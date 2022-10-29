package main

import (
	"fmt"

	"github.com/RafaelRochaS/journey-api/routes"
	"github.com/RafaelRochaS/journey-api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnvs()

	server := gin.New()
	server.SetTrustedProxies([]string{utils.TRUSTED_PROXIES})

	setMiddleware(server)
	routes.InitializeRoutes(server)

	server.Run(fmt.Sprintf("%s:%d", utils.HOST, utils.PORT))
}

func setMiddleware(server *gin.Engine) {
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
}
