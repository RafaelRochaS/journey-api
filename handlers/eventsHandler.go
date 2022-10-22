package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleEvents(rg *gin.RouterGroup) {
	events := rg.Group("/events")

	events.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "events")
	})

	events.POST("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, "done")
	})
}
