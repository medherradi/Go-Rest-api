package main

import (
	"RESTAPI/project/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // engine pointer

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run() // in dev it will be localhost:8080
}

func getEvents(ctx *gin.Context) {
	// we can use ctx.JSON(200, data)
	events := models.GetAllEvents()
	ctx.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvent(ctx *gin.Context) {

	// lets first say we have some data coming from the client side how can i access it
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		if strings.Contains(err.Error(), "Name") {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
			return
		} else if strings.Contains(err.Error(), "Description") {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Description is required"})
			return
		} else if strings.Contains(err.Error(), "Location") {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Location is required"})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "something went wrong please try again later"})
		}
		return
	}
	fmt.Println(event)
	ctx.JSON(http.StatusCreated, gin.H{"message": "event has been created successfully"})
}
