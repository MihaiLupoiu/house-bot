package models

// Flat description
type Flat struct {
	Title       string
	Description string
	Address     string
	Email       string
	Phone       string
	Price       int
	URL         string
	Picture     string
	Surface     int
	Floor       int
	Elevator    bool
	Rooms       int
	Bathrooms   int
}

// Coordinates of the flat
type Coordinates struct {
	Accuracy  int
	Latitude  float64
	Longitude float64
}
