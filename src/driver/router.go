package driver

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, handler *Handler) {
	api := r.Group("/api/drivers")
	{
		api.POST("/signup", handler.Signup)
		api.POST("/:id/location", handler.UpdateLocation)
	}

	ride := r.Group("/api/rides")
	{
		ride.POST("/:id/accept", handler.AcceptRide)
		ride.POST("/:id/complete", handler.CompleteRide)
	}
}
