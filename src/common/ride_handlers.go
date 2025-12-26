package common

import (
	"fmt"
	"sync"
	"time"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Driver struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
	Status   string   `json:"status"` // available, busy, offline
}

type Rider struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Ride struct {
	ID          string    `json:"id"`
	RiderID     string    `json:"rider_id"`
	DriverID    string    `json:"driver_id,omitempty"`
	Pickup      Location  `json:"pickup"`
	Dropoff     Location  `json:"dropoff"`
	Status      string    `json:"status"` // requested, accepted, in_progress, completed, cancelled
	RequestTime time.Time `json:"request_time"`
	AcceptTime  *time.Time `json:"accept_time,omitempty"`
	CompleteTime *time.Time `json:"complete_time,omitempty"`
	Fare        float64   `json:"fare,omitempty"`
}

type RideManager struct {
	drivers map[string]*Driver
	riders  map[string]*Rider
	rides   map[string]*Ride
	mu      sync.RWMutex
	nextID  int
}

func NewRideManager() *RideManager {
	return &RideManager{
		drivers: make(map[string]*Driver),
		riders:  make(map[string]*Rider),
		rides:   make(map[string]*Ride),
	}
}

// Rider Methods
func (rm *RideManager) CreateRider(r Rider) (*Rider, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	
	if _, exists := rm.riders[r.ID]; exists {
		return nil, fmt.Errorf("rider with ID %s already exists", r.ID)
	}
	
	rm.riders[r.ID] = &r
	return &r, nil
}

// Driver Methods
func (rm *RideManager) CreateDriver(d Driver) (*Driver, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	
	if _, exists := rm.drivers[d.ID]; exists {
		return nil, fmt.Errorf("driver with ID %s already exists", d.ID)
	}
	
	d.Status = "available"
	rm.drivers[d.ID] = &d
	return &d, nil
}

func (rm *RideManager) UpdateDriverLocation(driverID string, loc Location) (*Driver, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	
	driver, exists := rm.drivers[driverID]
	if !exists {
		return nil, fmt.Errorf("driver not found")
	}
	
	driver.Location = loc
	return driver, nil
}

// Ride Methods
func (rm *RideManager) CreateRideRequest(riderID string, pickup, dropoff Location) (*Ride, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	
	// Verify rider exists
	if _, exists := rm.riders[riderID]; !exists {
		return nil, fmt.Errorf("rider not found")
	}
	
	rm.nextID++
	ride := &Ride{
		ID:          fmt.Sprintf("ride_%d", rm.nextID),
		RiderID:     riderID,
		Pickup:      pickup,
		Dropoff:     dropoff,
		Status:      "requested",
		RequestTime: time.Now(),
		Fare:        rm.calculateFare(pickup, dropoff),
	}
	
	rm.rides[ride.ID] = ride
	
	// Find available drivers (simplified - just returns list)
	drivers := rm.findAvailableDrivers()
	if len(drivers) > 0 {
		// In a real system, you'd notify drivers here
		fmt.Printf("Notifying %d available drivers about ride %s\n", len(drivers), ride.ID)
	}
	
	return ride, nil
}

func (rm *RideManager) AcceptRide(rideID, driverID string) (*Ride, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	
	ride, exists := rm.rides[rideID]
	if !exists {
		return nil, fmt.Errorf("ride not found")
	}
	
	driver, exists := rm.drivers[driverID]
	if !exists {
		return nil, fmt.Errorf("driver not found")
	}
	
	if ride.Status != "requested" {
		return nil, fmt.Errorf("ride is not available for acceptance")
	}
	
	if driver.Status != "available" {
		return nil, fmt.Errorf("driver is not available")
	}
	
	// Update ride
	ride.DriverID = driverID
	ride.Status = "accepted"
	now := time.Now()
	ride.AcceptTime = &now
	
	// Update driver status
	driver.Status = "busy"
	
	return ride, nil
}

func (rm *RideManager) CompleteRide(rideID, driverID string) (*Ride, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	
	ride, exists := rm.rides[rideID]
	if !exists {
		return nil, fmt.Errorf("ride not found")
	}
	
	if ride.DriverID != driverID {
		return nil, fmt.Errorf("driver is not assigned to this ride")
	}
	
	if ride.Status != "accepted" && ride.Status != "in_progress" {
		return nil, fmt.Errorf("ride cannot be completed in current status")
	}
	
	// Update ride
	ride.Status = "completed"
	now := time.Now()
	ride.CompleteTime = &now
	
	// Update driver status back to available
	if driver, exists := rm.drivers[driverID]; exists {
		driver.Status = "available"
	}
	
	return ride, nil
}

func (rm *RideManager) CancelRide(rideID string) (*Ride, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	
	ride, exists := rm.rides[rideID]
	if !exists {
		return nil, fmt.Errorf("ride not found")
	}
	
	if ride.Status == "completed" {
		return nil, fmt.Errorf("cannot cancel completed ride")
	}
	
	// If a driver was assigned, make them available again
	if ride.DriverID != "" {
		if driver, exists := rm.drivers[ride.DriverID]; exists {
			driver.Status = "available"
		}
	}
	
	ride.Status = "cancelled"
	return ride, nil
}

func (rm *RideManager) GetRide(rideID string) (*Ride, error) {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	
	ride, exists := rm.rides[rideID]
	if !exists {
		return nil, fmt.Errorf("ride not found")
	}
	return ride, nil
}

// Helper methods
func (rm *RideManager) findAvailableDrivers() []*Driver {
	var available []*Driver
	for _, driver := range rm.drivers {
		if driver.Status == "available" {
			available = append(available, driver)
		}
	}
	return available
}

func (rm *RideManager) calculateFare(pickup, dropoff Location) float64 {
	// Simplified fare calculation
	baseFare := 5.0
	// In a real system, calculate distance and multiply by rate
	distance := 5.0 // Mock distance in km
	ratePerKm := 2.0
	return baseFare + (distance * ratePerKm)
}
// Add these methods to your common/ride_manager.go file:

// GetAvailableRides returns all rides that are waiting for a driver
func (rm *RideManager) GetAvailableRides() []*Ride {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	
	var available []*Ride
	for _, ride := range rm.rides {
		if ride.Status == "requested" {
			available = append(available, ride)
		}
	}
	return available
}
func (rm *RideManager) GetAvailableDrivers() []*Driver {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	
	var available []*Driver
	for _, driver := range rm.drivers {
		if driver.Status == "available" {
			available = append(available, driver)
		}
	}
	return available
}
// GetDriverRides returns all rides for a specific driver
func (rm *RideManager) GetDriverRides(driverID string) []*Ride {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	
	var driverRides []*Ride
	for _, ride := range rm.rides {
		if ride.DriverID == driverID {
			driverRides = append(driverRides, ride)
		}
	}
	return driverRides
}

// GetRiderRides returns all rides for a specific rider
func (rm *RideManager) GetRiderRides(riderID string) []*Ride {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	
	var riderRides []*Ride
	for _, ride := range rm.rides {
		if ride.RiderID == riderID {
			riderRides = append(riderRides, ride)
		}
	}
	return riderRides
}
