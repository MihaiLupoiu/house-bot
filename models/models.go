package models

type Telegram struct {
	BotID        string
	ChannelID    string
	SendMessages bool
}
type Filters struct {
	LocationName     string `js:"-" json:"LocationName"`
	MaximumPrice     int    `js:"-" json:"Maximum_price"`
	MinimumPrice     int    `js:"-" json:"Minimum_price"`
	MaximumSurface   int    `js:"-" json:"Maximum_sueface"`
	MinimumSurface   int    `js:"-" json:"Minimum_surface"`
	MaximumRooms     int    `js:"-" json:"Maximum_rooms"`
	MinimumRooms     int    `js:"-" json:"Minimum_rooms"`
	MinimumBathrooms int    `js:"-" json:"Minimum_bathrooms"`
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
	Debug    bool
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
