package routes

import (
	"bookcabin-seatmap-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterSegmentRoutes(r *gin.Engine) {
	r.GET("/api/segment", controllers.GetSegment)
}
