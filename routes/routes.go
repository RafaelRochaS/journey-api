package routes

import (
	"github.com/RafaelRochaS/journey-api/handlers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(server *gin.Engine) {
	v1 := server.Group("/v1")

	handlers.HandleHealthcheck(v1)
	handlers.HandleEvents(v1)
}
