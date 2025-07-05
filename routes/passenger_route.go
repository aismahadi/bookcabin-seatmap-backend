package routes

import (
	"bookcabin-seatmap-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPassengerRoutes(r *gin.Engine) {
	r.GET("/api/passenger", controllers.GetPassenger)
}
