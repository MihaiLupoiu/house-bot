package fotocasa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MihaiLupoiu/house-bot/lib/util"
	"github.com/MihaiLupoiu/house-bot/models"
	bolt "github.com/coreos/bbolt"
)

// Location ...
type Location struct {
	Location string `js:"-" json:"location"`
	Zone     string `js:"-" json:"zone"`
	Text     string `js:"-" json:"text"`
}

// LocationSegments ...
type LocationSegments struct {
	IDs             string             `js:"-" json:"ids"`
	Coordinates     models.Coordinates `js:"-" json:"coordinates"`
	LocationLiteral string             `js:"-" json:"locationLiteral"`
}

type Detail struct {
	IDs             string             `js:"-" json:"es"`
	Coordinates     models.Coordinates `js:"-" json:"coordinates"`
	LocationLiteral string             `js:"-" json:"locationLiteral"`
}

type Transactions struct {
	TypeID  int     `js:"-" json:"transactionTypeId"`
	Value   []int64 `js:"-" json:"value"`
	Reduced int64   `js:"-" json:"reduced"`
}

type Address struct {
	Ubication   string             `js:"-" json:"ubication"`
	Location    map[string]string  `js:"-" json:"location"`
	Coordinates models.Coordinates `js:"-" json:"coordinates"`
}

type HouseDetail struct {
	ES string `js:"-" json:"es"`
}

type Features struct {
	Key   string `js:"-" json:"key"`
	Value []int  `js:"-" json:"value"`
}

type FotocasaHouse struct {
	ID            int64                  `js:"-" json:"id"`
	New           bool                   `js:"-" json:"isNew"`
	Premium       bool                   `js:"-" json:"isPremium"`
	Advertiser    map[string]interface{} `js:"-" json:"advertiser"`
	Detail        HouseDetail            `js:"-" json:"detail"`
	Multimedias   []interface{}          `js:"-" json:"multimedias"`
	Features      []Features             `js:"-" json:"features"`
	OtherFeatures []interface{}          `js:"-" json:"otherFeatures"`
	Transactions  []Transactions         `js:"-" json:"transactions"`
	Products      []interface{}          `js:"-" json:"products"`
	Address       Address                `js:"-" json:"address"`
	Date          time.Time              `js:"-" json:"date"`
	Description   string                 `js:"-" json:"description"`
}

type FotocasaHouses struct {
	Count  int64           `js:"-" json:"count"`
	Houses []FotocasaHouse `js:"-" json:"realEstates"`
}

var (
	apiURL    = "https://api.fotocasa.es/PropertySearch/"
	webURL    = "https://www.fotocasa.es"
	db        *bolt.DB
	config    models.Fotocasa
	filters   models.Filters
	HouseChan chan *models.House

	// FotocasaLocas
	transactionTypeId = 1 // buy by default
	locSegments       LocationSegments
)

// Init Fotocasa
func Init(database *bolt.DB, communicationChannel chan *models.House, configuration models.Fotocasa, searchFilters models.Filters) {
	db = database
	HouseChan = communicationChannel
	config = configuration
	filters = searchFilters

	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("fotocasa"))
		return err
	})
	if err != nil {
		log.Println(err)
		return
	}
	switch config.TransactionType {
	case "Buy":
		transactionTypeId = 2
	case "Rent":
		transactionTypeId = 5
	default:
		// Buy
		transactionTypeId = 1
	}

	if config.SortType != "bumpdate" && config.SortType != "publicationDate" {
		log.Println("Invalid Fotocasa.ShortType, changed to default 'bumpdate'.")
		config.SortType = "bumpdate"
	}

	// TODO: Fix. Sometimes only CombinedLocationIds is necesarry. Latidute and Longitude not always.
	if config.CombinedLocationIds != "" && config.Latitude != 0 && config.Longitude != 0 {
		locSegments.IDs = config.CombinedLocationIds
		locSegments.Coordinates = models.Coordinates{
			Accuracy:  1,
			Latitude:  config.Latitude,
			Longitude: config.Longitude,
		}
		log.Printf("Found valid config for Fotocasa.combinedLocationID. Values: %v\n", locSegments)
	} else {

		if filters.LocationName == "" {
			filters.LocationName = "Castellon"
		}
		location, err := getCityLocation(filters.LocationName)
		if err != nil {
			log.Println(err)
			return
		}
		locSegments, err = getCombinedLocationIds(location)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Got Fotocasa.combinedLocationID from '%s' Values: %v\n", filters.LocationName, locSegments)
	}

}

func getCityLocation(query string) (Location, error) {
	var l Location
	searchQuery := apiURL + "SearchUrl?latitude=0&longitude=0&text=" + query + "&culture=es-ES"

	body, err := util.Get(searchQuery)
	if err != nil {
		log.Println("error:", err)
		return l, err
	}

	err = json.Unmarshal(body, &l)
	if err != nil {
		log.Printf("error: %s: %s", err, body)
		return l, err
	}
	return l, nil
}

func getCombinedLocationIds(location Location) (LocationSegments, error) {
	var ls LocationSegments

	searchQuery := apiURL + "UrlLocationSegments?location=" + location.Location + "&zone=" + location.Zone
	body, err := util.Get(searchQuery)
	if err != nil {
		log.Println("error:", err)
		return ls, err
	}

	err = json.Unmarshal(body, &ls)
	if err != nil {
		log.Printf("error: %s: %s", err, body)
		return ls, err
	}
	return ls, nil
}

func getCityDetailsFromCoordinates(coord models.Coordinates) interface{} {
	var repl interface{}
	searchQuery := "https://reverse.geocoder.cit.api.here.com/6.2/reversegeocode.json?prox="
	searchQuery += strconv.FormatFloat(coord.Latitude, 'f', -1, 64)
	searchQuery += "%2C" + strconv.FormatFloat(coord.Longitude, 'f', -1, 64)
	searchQuery += "&mode=retrieveAddresses&maxresults=1&gen=8&app_id=DemoAppId01082013GAL&app_code=AJKnXv84fjrb0KIHawS0Tg"

	body, err := util.Get(searchQuery)
	if err != nil {
		log.Println("error:", err)
		return repl
	}

	err = json.Unmarshal(body, &repl)
	if err != nil {
		log.Printf("error: %s: %s", err, body)
		return repl
	}
	return repl
}

func getHouses(ls LocationSegments, page int) []FotocasaHouse {
	var h FotocasaHouses
	searchQuery := apiURL + "Search?"
	searchQuery += "combinedLocationIds=" + ls.IDs
	searchQuery += "&culture=es-ES"
	searchQuery += "&hrefLangCultures=ca-ES%3Bes-ES%3Bde-DE%3Ben-GB"
	searchQuery += "&isNewConstruction=false"
	searchQuery += "&isMap=false"
	searchQuery += "&latitude=" + strconv.FormatFloat(ls.Coordinates.Latitude, 'f', -1, 64)
	searchQuery += "&longitude=" + strconv.FormatFloat(ls.Coordinates.Longitude, 'f', -1, 64)
	searchQuery += "&pageNumber=" + strconv.Itoa(page)
	searchQuery += "&propertyTypeId=2"
	searchQuery += "&sortOrderDesc=true"
	searchQuery += "&sortType=" + config.SortType
	searchQuery += "&transactionTypeId=" + strconv.Itoa(transactionTypeId)

	body, err := util.Get(searchQuery)
	if err != nil {
		log.Println("error:", err)
	}

	err = json.Unmarshal(body, &h)
	if err != nil {
		log.Printf("error: %s: %s", err, body)
	}

	return h.Houses
}

func getAllHouses(ls LocationSegments) []FotocasaHouse {
	var houses []FotocasaHouse
	page := 1
	for {
		recvHouses := getHouses(ls, page)
		if len(recvHouses) > 0 {
			houses = append(houses, recvHouses...)
			page++
		} else {
			break
		}
	}
	return houses

}

func processNewHouse(house FotocasaHouse) error {
	return processHouse(house, "new")
}

func processHouseNewPrice(house FotocasaHouse) error {
	return processHouse(house, "changed")
}

func processHouse(house FotocasaHouse, stastus string) error {
	houseID := strconv.FormatInt(house.ID, 10)
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("fotocasa"))
		encoded, err := json.Marshal(house)
		if err != nil {
			return err
		}

		b.Put([]byte(houseID), encoded)
		return nil
	})
	if err != nil {
		return err
	}

	addr := house.Address.Location["level5"] + " " + house.Address.Location["level6"] + " " + house.Address.Location["level7"] + " " + house.Address.Location["level8"]

	surface, rooms, bathrooms := -1, -1, -1
	for _, feature := range house.Features {
		switch feature.Key {
		case "rooms":
			rooms = feature.Value[0]
		case "bathrooms":
			bathrooms = feature.Value[0]
		case "surface":
			surface = feature.Value[0]
		default:
			log.Printf("Unkown features %v\n", feature)
		}
	}

	var title string
	if stastus == "new" {
		title = "New House ID:"
	} else {
		title = "Changed House Price House ID:"
	}

	//TODO: add pictures
	HouseChan <- &models.House{
		Title:       title + houseID,
		Price:       int(house.Transactions[0].Value[0]),
		Reduced:     int(house.Transactions[0].Reduced),
		URL:         fmt.Sprintf("https://www.fotocasa.es%s", house.Detail.ES),
		Description: house.Description,
		Address:     addr,
		Surface:     surface,
		Rooms:       rooms,
		Bathrooms:   bathrooms,
		// Picture:
	}

	return nil
}

// TickerCheck ...
func TickerCheck(ctx context.Context, ctxCancel context.CancelFunc) {
	defer ctxCancel()

	interval := time.Minute * time.Duration(config.MinutesCheckInterval)
	t := time.NewTimer(time.Second)
	for {
		select {
		case <-ctx.Done():
			return

		case <-t.C:
			log.Println("Time to search!")
			skipped := 0
			page := 1
			for {
				houses := getHouses(locSegments, page)
				if len(houses) > 0 {
					// Check if house already in db or price changed.
					for _, house := range houses {
						houseID := strconv.FormatInt(house.ID, 10)

						err := db.View(func(tx *bolt.Tx) error {
							b := tx.Bucket([]byte("fotocasa"))
							v := b.Get([]byte(houseID))
							if v == nil {
								log.Println("house id:" + houseID + " dosn't exists")
								return nil
							}

							var h FotocasaHouse
							if err := json.Unmarshal(v, &h); err != nil {
								log.Printf("error: %s: %s\n", err, v)
							}

							// Cheking if price changed
							if house.Transactions[0].Value[0] == h.Transactions[0].Value[0] {
								return errors.New("house exists")
							} else {
								log.Printf("House price changed!!!!!!\n")
								return errors.New("price changed")
							}
						})
						if err != nil {
							if err == errors.New("price changed") {
								processHouseNewPrice(house)
							} else {
								skipped++
							}
							continue
						}
						processNewHouse(house)
					}
				} else {
					// end of list
					break
				}

				if skipped > 50 {
					// If skipped 50 houses stop searching. No more new houses.
					break
				}

				page++
			}

			if skipped > 0 {
				log.Printf("%d fotocasas skipped\n", skipped)
			}
			t.Reset(interval)
		}
	}

}
