package driver

import "fmt"

type Service struct {
	store map[string]Driver
}

func NewService() *Service {
	return &Service{store: make(map[string]Driver)}
}

func (s *Service) CreateDriver(d Driver) Driver {
	s.store[d.ID] = d
	return d
}

func (s *Service) UpdateLocation(id, location string) (Driver, error) {
	d, exists := s.store[id]
	if !exists {
		return Driver{}, fmt.Errorf("driver not found")
	}
	d.Location = location
	s.store[id] = d
	return d, nil
}