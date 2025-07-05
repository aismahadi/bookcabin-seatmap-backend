package controllers

import (
	"bookcabin-seatmap-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSegment(c *gin.Context) {
	segment, err := services.GetSegment()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Segment not found"})
		return
	}
	c.JSON(http.StatusOK, segment)
}
