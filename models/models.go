package models

type Telegram struct {
	BotID     string
	ChannelID string
}
type Filters struct {
	LocationName     string
	MaximumPrice     int
	MinimumPrice     int
	MaximumRooms     int
	MinimumRooms     int
	MinimumBathrooms int
	MinimumPhotos    int
}
type Fotocasa struct {
	TransactionType      string
	SortType             string
	CombinedLocationIds  string
	Latitude             float64
	Longitude            float64
	MinutesCheckInterval int
}
type Config struct {
	Telegram Telegram
	Database string
	Filters  Filters
	Fotocasa Fotocasa
}

// House description
type House struct {
	Title       string
	Description string
	Address     string
	Email       string
	Phone       string
	Price       int
	Reduced     int
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
