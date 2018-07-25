package fotocasa

import (
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

type FotocasaHouse struct {
	ID            int64                  `js:"-" json:"id"`
	New           bool                   `js:"-" json:"isNew"`
	Premium       bool                   `js:"-" json:"isPremium"`
	Advertiser    map[string]interface{} `js:"-" json:"advertiser"`
	Detail        map[string]interface{} `js:"-" json:"detail"`
	Multimedias   []interface{}          `js:"-" json:"multimedias"`
	Features      []interface{}          `js:"-" json:"features"`
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
	apiURL = "https://api.fotocasa.es/PropertySearch/"
	webURL = "https://www.fotocasa.es"
)

func GetCityLocation(query string) (Location, error) {
	var l Location
	searchQuery := apiURL + "SearchUrl?latitude=0&longitude=0&text=" + query + "&culture=es-ES"
	body, err := util.Get(searchQuery)
	if err != nil {
		log.Println("error:", err)
		return l, err
	}

	err = json.Unmarshal(body, &l)
	if err != nil {
		fmt.Printf("error: %s: %s", err, body)
		return l, err
	}
	return l, nil
}

func GetCombinedLocationIds(location Location) (LocationSegments, error) {
	var ls LocationSegments

	searchQuery := apiURL + "UrlLocationSegments?location=" + location.Location + "&zone=" + location.Zone
	body, err := util.Get(searchQuery)
	if err != nil {
		log.Println("error:", err)
		return ls, err
	}

	err = json.Unmarshal(body, &ls)
	if err != nil {
		fmt.Printf("error: %s: %s", err, body)
		return ls, err
	}
	return ls, nil
}

func GetHouses(ls LocationSegments, page int) []FotocasaHouse {
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
	searchQuery += "&sortType=bumpdate"
	searchQuery += "&transactionTypeId=1"

	// transactionTypeId:
	// Comprar 1
	// Alquiler 5

	// sortType
	// bumpdate
	// publicationDate

	body, err := util.Get(searchQuery)
	if err != nil {
		log.Println("error:", err)
	}

	err = json.Unmarshal(body, &h)
	if err != nil {
		fmt.Printf("error: %s: %s", err, body)
	}

	/*
		for i, v := range h.Houses {
			fmt.Printf("-->> %d %v %v\n", i, v.ID, v.Address)
		}

	*/

	// fmt.Printf("-->> %v\n", h.Houses[0])
	return h.Houses
}

func getAllHouses(ls LocationSegments) []FotocasaHouse {

	var houses []FotocasaHouse

	page := 1
	for {
		recvHouses := GetHouses(ls, page)
		if len(recvHouses) > 0 {
			houses = append(houses, recvHouses...)
			page++
			fmt.Println(page)
		} else {
			break
		}

		if page == 5 {
			break
		}
	}
	return houses

}

func TickerCheck(db *bolt.DB) {

	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("fotocasa"))
		return err
	})
	if err != nil {
		log.Println(err)
		return
	}

	//interval := time.Minute * 30
	//interval := time.Second * 5 //DEBUG
	t := time.NewTimer(time.Second)
	for {
		select {
		/*
			case <-ctx.Done():
				return
		*/
		case <-t.C:

			location, _ := GetCityLocation("castellon")
			locDet, _ := GetCombinedLocationIds(location)

			// TODO: If skip > 100 stop going in more pages
			houses := getAllHouses(locDet)

			skipped := 0
			for _, house := range houses {

				houseID := strconv.FormatInt(house.ID, 10)

				err := db.View(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("fotocasa"))
					v := b.Get([]byte(houseID))
					if v == nil {
						log.Println("flat dosn't exists")
						return nil
					}

					var h FotocasaHouse
					if err = json.Unmarshal(v, &h); err != nil {
						fmt.Printf("error: %s: %s", err, v)
					}

					//fmt.Println("FOUND ->", h.ID, "Adress:", h.Address, "Transaction:", h.Transactions)
					//fmt.Println("COMRD ->", house.ID, "Adress:", house.Address, "Transaction:", house.Transactions)
					// Cheking if price changed
					if house.Transactions[0].Value[0] == h.Transactions[0].Value[0] {
						return errors.New("flat exists")
					}

					return nil
				})
				if err != nil {
					fmt.Println("Skipped: ", err)
					skipped++
					continue
				}

				err = db.Update(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("fotocasa"))

					//fmt.Println("ADDING->", house.ID, "Adress:", house.Address, "Transaction:", house.Transactions)
					encoded, err := json.Marshal(house)
					if err != nil {
						return err
					}

					b.Put([]byte(houseID), encoded)
					return nil
				})
				if err != nil {
					continue
				}

			}

			if skipped > 0 {
				log.Printf("%d fotocasas skipped\n", skipped)
			}
			//t.Reset(interval)
		}
	}

}
