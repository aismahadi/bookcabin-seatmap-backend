package controllers

import (
	"bookcabin-seatmap-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSeatMap(c *gin.Context) {
	aircrafts := services.GetFullSeatMap()
	c.JSON(http.StatusOK, aircrafts)
}
