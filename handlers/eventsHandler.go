package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RafaelRochaS/journey-api/models"
	"github.com/RafaelRochaS/journey-api/producers"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
)

func HandleEvents(rg *gin.RouterGroup) {
	producer := setUpProducer()

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

		eventObjectDto := &models.EventObjectDto{
			Event:     event,
			Timestamp: time.Now().In(time.UTC),
			TrxId:     shortuuid.New(),
		}

		log.Println("Received event: ")
		log.Println("Sensor type: ", event.SensorType)
		log.Println("Measure 1: ", event.Measure1)
		log.Println("Measure 2: ", event.Measure2)
		log.Println("Measure 3: ", event.Measure3)
		log.Println("Measure 4: ", event.Measure4)
		log.Println("Time: ", eventObjectDto.Timestamp)
		log.Println("TrxID: ", eventObjectDto.TrxId)

		err := producer.HandleEventProduction(*eventObjectDto)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Failure to set up event. Please wait a little and try again\nError: %s", err.Error()))
			return
		}

		ctx.JSON(http.StatusAccepted, "Event received. Added to processing queue.")
	})
}

func setUpProducer() producers.EventProducer {
	producer, err := producers.NewKafkaProducer()

	if err != nil {
		log.Panic("Failure to set up event producer! ", err.Error())
	}

	return producer
}
