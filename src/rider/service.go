package rider

import (
	"fmt"
	"sync"
)

type Service struct {
	riders   map[string]Rider
	rides    map[string]RideRequest
	mu       sync.RWMutex
	nextID   int
}

func NewService() *Service {
	return &Service{
		riders: make(map[string]Rider),
		rides:  make(map[string]RideRequest),
	}
}

func (s *Service) CreateRider(r Rider) Rider {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.riders[r.ID] = r
	return r
}

func (s *Service) CreateRideRequest(req RideRequestInput) RideRequest {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.nextID++
	ride := RideRequest{
		ID:      fmt.Sprintf("ride_%d", s.nextID),
		RiderID: req.RiderID,
		Pickup:  req.Pickup,
		Dropoff: req.Dropoff,
		Status:  "requested",
	}
	s.rides[ride.ID] = ride
	return ride
}

func (s *Service) GetRideStatus(id string) (RideRequest, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	ride, exists := s.rides[id]
	if !exists {
		return RideRequest{}, fmt.Errorf("ride not found")
	}
	return ride, nil
}

func (s *Service) CancelRide(id string) (RideRequest, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	ride, exists := s.rides[id]
	if !exists {
		return RideRequest{}, fmt.Errorf("ride not found")
	}
	
	if ride.Status == "completed" {
		return RideRequest{}, fmt.Errorf("cannot cancel completed ride")
	}
	
	ride.Status = "cancelled"
	s.rides[id] = ride
	return ride, nil
}