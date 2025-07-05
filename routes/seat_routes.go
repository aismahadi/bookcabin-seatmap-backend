package routes

import (
	"bookcabin-seatmap-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterSeatRoutes(r *gin.Engine) {
	r.GET("/api/seatmap", controllers.GetSeatMap)
}
