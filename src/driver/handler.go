package driver

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
	var d Driver
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := h.service.CreateDriver(d)
	c.JSON(http.StatusOK, gin.H{"message": "Driver created", "driver": created})
}

func (h *Handler) UpdateLocation(c *gin.Context) {
	id := c.Param("id")
	var loc struct {
		Location string `json:"location"`
	}
	if err := c.ShouldBindJSON(&loc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.service.UpdateLocation(id, loc.Location)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Location updated", "driver": updated})
}

func (h *Handler) AcceptRide(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Ride accepted", "ride_id": c.Param("id")})
}

func (h *Handler) CompleteRide(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Ride completed", "ride_id": c.Param("id")})
}
