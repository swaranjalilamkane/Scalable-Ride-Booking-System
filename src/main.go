package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"GroupProject/common"
	"os"
	"encoding/json"
)

type App struct {
	rideManager *common.RideManager
}

func NewApp() *App {
	app := &App{
		rideManager: common.NewRideManager(),
	}

	// Load 50 drivers from JSON file
	data, err := os.ReadFile("drivers.json")
	if err != nil {
		panic("failed to load drivers.json: " + err.Error())
	}

	var drivers []common.Driver
	if err := json.Unmarshal(data, &drivers); err != nil {
		panic("invalid drivers.json: " + err.Error())
	}

	for _, d := range drivers {
		app.rideManager.CreateDriver(d)
	}

	return app
}

// Rider Handlers
func (app *App) RiderSignup(c *gin.Context) {
	var rider common.Rider
	if err := c.ShouldBindJSON(&rider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	created, err := app.rideManager.CreateRider(rider)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Rider created", "rider": created})
}

func (app *App) RequestRide(c *gin.Context) {
	var req struct {
		RiderID string          `json:"rider_id"`
		Pickup  common.Location `json:"pickup"`
		Dropoff common.Location `json:"dropoff"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ride, err := app.rideManager.CreateRideRequest(req.RiderID, req.Pickup, req.Dropoff)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Ride requested", "ride": ride})
}

func (app *App) GetRideStatus(c *gin.Context) {
	rideID := c.Param("id")
	
	ride, err := app.rideManager.GetRide(rideID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"ride": ride})
}

func (app *App) CancelRide(c *gin.Context) {
	rideID := c.Param("id")
	
	ride, err := app.rideManager.CancelRide(rideID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Ride cancelled", "ride": ride})
}

// Driver Handlers
func (app *App) DriverSignup(c *gin.Context) {
	var driver common.Driver
	if err := c.ShouldBindJSON(&driver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	created, err := app.rideManager.CreateDriver(driver)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Driver created", "driver": created})
}

func (app *App) UpdateDriverLocation(c *gin.Context) {
	driverID := c.Param("id")
	
	var loc struct {
		Location common.Location `json:"location"`
	}
	
	if err := c.ShouldBindJSON(&loc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	driver, err := app.rideManager.UpdateDriverLocation(driverID, loc.Location)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Location updated", "driver": driver})
}

func (app *App) AcceptRide(c *gin.Context) {
	rideID := c.Param("id")
	
	var req struct {
		DriverID string `json:"driver_id"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ride, err := app.rideManager.AcceptRide(rideID, req.DriverID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Ride accepted", "ride": ride})
}

func (app *App) CompleteRide(c *gin.Context) {
	rideID := c.Param("id")
	
	var req struct {
		DriverID string `json:"driver_id"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ride, err := app.rideManager.CompleteRide(rideID, req.DriverID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Ride completed", "ride": ride})
}
func (app *App) GetAvailableDrivers(c *gin.Context) {
	drivers := app.rideManager.GetAvailableDrivers()
	c.JSON(http.StatusOK, gin.H{"drivers": drivers})
}
func (app *App) GetAvailableRides(c *gin.Context) {
	rides := app.rideManager.GetAvailableRides()
	c.JSON(http.StatusOK, gin.H{"rides": rides})
}

func (app *App) GetDriverRides(c *gin.Context) {
	driverID := c.Param("id")
	rides := app.rideManager.GetDriverRides(driverID)
	c.JSON(http.StatusOK, gin.H{"rides": rides})
}

func (app *App) GetRiderHistory(c *gin.Context) {
	riderID := c.Param("id")
	rides := app.rideManager.GetRiderRides(riderID)
	c.JSON(http.StatusOK, gin.H{"rides": rides})
}

func main() {
	r := gin.Default()
	app := NewApp()
	
	// Rider routes
	rides := r.Group("/api/rides")
	{
		rides.POST("/signup", app.RiderSignup)
		rides.POST("/request", app.RequestRide)
		rides.GET("/:id/status", app.GetRideStatus)
		rides.POST("/:id/cancel", app.CancelRide)
	}
	
	// Driver routes
	drivers := r.Group("/api/drivers")
	{
		drivers.POST("/signup", app.DriverSignup)
		drivers.POST("/:id/location", app.UpdateDriverLocation)
	}
	
	// Shared ride routes (used by drivers)
	rideActions := r.Group("/api/rides")
	{
		rideActions.POST("/:id/accept", app.AcceptRide)
		rideActions.POST("/:id/complete", app.CompleteRide)
	}

	// New routes
	drivers.GET("/:id/rides", app.GetDriverRides)
	drivers.GET("/available-rides", app.GetAvailableRides)
	rides.GET("/rider/:id/history", app.GetRiderHistory)
	drivers.GET("/available-drivers", app.GetAvailableDrivers)
	
	common.Info("Starting Unified Ride Booking Service on port 8080...")
	r.Run(":8080")
}