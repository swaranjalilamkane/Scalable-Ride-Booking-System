package rider

type Rider struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

type RideRequest struct {
	ID       string `json:"id"`
	RiderID  string `json:"rider_id"`
	DriverID string `json:"driver_id,omitempty"`
	Pickup   string `json:"pickup"`
	Dropoff  string `json:"dropoff"`
	Status   string `json:"status"` // requested, accepted, completed, cancelled
}

type RideRequestInput struct {
	RiderID string `json:"rider_id"`
	Pickup  string `json:"pickup"`
	Dropoff string `json:"dropoff"`
}