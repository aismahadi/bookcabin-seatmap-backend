package controllers

import (
	"bookcabin-seatmap-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSegment(c *gin.Context) {
	equipment := c.Param("equipment")
	segment, err := services.GetSegment(equipment)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Segment not found"})
		return
	}
	c.JSON(http.StatusOK, segment)
}
