package rider

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, handler *Handler) {
	api := r.Group("/api/rides")
	{
		api.POST("/signup", handler.Signup)
		api.POST("/request", handler.RequestRide)
		api.GET("/:id/status", handler.GetRideStatus)
		api.POST("/:id/cancel", handler.CancelRide)
	}
}