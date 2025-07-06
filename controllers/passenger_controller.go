package controllers

import (
	"bookcabin-seatmap-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPassenger(c *gin.Context) {
	id := c.Param("id")
	passenger, err := services.GetPassenger(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Passenger not found"})
		return
	}
	c.JSON(http.StatusOK, passenger)
}
