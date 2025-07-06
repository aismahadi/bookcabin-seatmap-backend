package controllers

import (
	"bookcabin-seatmap-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSeatMap(c *gin.Context) {
	id := c.Param("id")
	seatmap, err := services.GetSeatMap(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SeatMap not found"})
		return
	}
	c.JSON(http.StatusOK, seatmap)
}
