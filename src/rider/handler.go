package rider

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Signup(c *gin.Context) {
	var r Rider
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	created := h.service.CreateRider(r)
	c.JSON(http.StatusOK, gin.H{"message": "Rider created", "rider": created})
}

func (h *Handler) RequestRide(c *gin.Context) {
	var req RideRequestInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.RiderID == "" || req.Pickup == "" || req.Dropoff == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rider_id, pickup, and dropoff are required"})
		return
	}

	// -----------------------------
	// Check if rider already has a pending ride
	// -----------------------------
	activeRide, err := h.service.GetActiveRideForRider(req.RiderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if activeRide != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Rider already has an active ride. Complete it before requesting a new one.",
			"ride":  activeRide,
		})
		return
	}

	// Create new ride request
	ride := h.service.CreateRideRequest(req)
	c.JSON(http.StatusOK, gin.H{"message": "Ride requested", "ride": ride})
}

func (h *Handler) GetRideStatus(c *gin.Context) {
	id := c.Param("id")
	
	ride, err := h.service.GetRideStatus(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"ride": ride})
}

func (h *Handler) CancelRide(c *gin.Context) {
	id := c.Param("id")
	
	ride, err := h.service.CancelRide(id)
	if err != nil {
		if err.Error() == "ride not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Ride cancelled", "ride": ride})
}