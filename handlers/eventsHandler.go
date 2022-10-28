package handlers

import (
	"log"
	"net/http"

	"github.com/RafaelRochaS/journey-api/models"
	"github.com/gin-gonic/gin"
)

func HandleEvents(rg *gin.RouterGroup) {
	events := rg.Group("/events")

	events.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "[WIP] events")
	})

	events.POST("/", func(ctx *gin.Context) {
		var event models.EventObject

		if err := ctx.ShouldBindJSON(&event); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		log.Println("Received event: ")
		log.Println("Sensor type: ", event.SensorType)
		log.Println("Measure 1: ", event.Measure1)
		log.Println("Measure 2: ", event.Measure2)
		log.Println("Measure 3: ", event.Measure3)
		log.Println("Measure 4: ", event.Measure4)

		ctx.JSON(http.StatusAccepted, "Event received. Added to processing queue.")
	})
}
