package fotocasa

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MihaiLupoiu/house-bot/lib/util"
	"github.com/MihaiLupoiu/house-bot/models"
)

// Location ...
type Location struct {
	Location string `js:"-" json:"location"`
	Zone     string `js:"-" json:"zone"`
	Text     string `js:"-" json:"text"`
}

// LocationSegments ...
type LocationSegments struct {
	IDs             string             `js:"-" json:"location"`
	Coordinates     models.Coordinates `js:"-" json:"coordinates"`
	LocationLiteral string             `js:"-" json:"locationLiteral"`
}

type Detail struct {
	IDs             string             `js:"-" json:"es"`
	Coordinates     models.Coordinates `js:"-" json:"coordinates"`
	LocationLiteral string             `js:"-" json:"locationLiteral"`
}

type Transactions struct {
	TypeId  int     `js:"-" json:"transactionTypeId"`
	Value   []int64 `js:"-" json:"value"`
	Reduced int64   `js:"-" json:"reduced"`
}

type Address struct {
	Ubication   string             `js:"-" json:"ubication"`
	Location    map[string]string  `js:"-" json:"location"`
	Coordinates models.Coordinates `js:"-" json:"coordinates"`
}

type FotocasaHouse struct {
	ID            int64 `js:"-" json:"id"`
	New           bool
	Premium       bool
	Advertiser    map[string]interface{}
	Detail        map[string]interface{}
	Multimedias   map[interface{}]interface{}
	Features      map[string]interface{}
	OtherFeatures map[string]interface{}
	Transactions  []Transactions
	Products      map[string]interface{}
	Address       Address
	Date          time.Time
	Description   string
}

/*
{
	"advertiser": {
		"logo": {
			"multimedia": {
				"url": "https://d.fotocasa.es/client/9202753963841/361703.jpg/100x52/",
				"typeId": 8
			},
			"url": {
				"es": "/inmobiliarias/haya-real-estate-9202753963841"
			}
		},
		"phone": "",
		"typeId": 3,
		"clientId": 9202753963841
	},
	"detail": {
		"es": "/vivienda/alcala-de-xivert/alcossebre-aire-acondicionado-parking-terraza-piscina-alborcer-5-147469155"
	},
	"multimedias": [
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428896933.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428896949.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428896963.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428896977.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428896991.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897002.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897016.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897041.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897058.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897078.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897097.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897117.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897135.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897152.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897178.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897198.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897213.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897228.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897249.jpg",
			"typeId": 2
		},
		{
			"url": "https://d.fotocasa.es/anuncio/2018/06/28/147469155/428897268.jpg",
			"typeId": 2
		}
	],
	"features": [
		{
			"key": "rooms",
			"value": [
				4
			]
		},
		{
			"key": "bathrooms",
			"value": [
				4
			]
		},
		{
			"key": "surface",
			"value": [
				206
			]
		}
	],
	"otherFeatures": null,
	"transactions": [
		{
			"transactionTypeId": 1,
			"value": [
				224700
			],
			"reduced": 0
		}
	],
	"products": [],
	"address": {
		"ubication": "Alborcer, 5",
		"location": {
			"country": "España",
			"level1": "Comunitat Valenciana",
			"level2": "Castellón",
			"level3": "Baix Maestrat",
			"level4": "Alcalà de Xivert, Zona de",
			"level5": "Alcalà de Xivert",
			"level6": "Alcossebre",
			"level7": "",
			"level8": "",
			"upperLevel": "Alcossebre"
		},
		"coordinates": {
			"accuracy": 1,
			"latitude": 40.22890163,
			"longitude": 0.2686177
		}
	},
	"date": "2018-07-22T13:33:09.5151974Z",
	"description": "Chalet adosado de cuatro alturas compuesta por cuatro dormitorios y tres baños en la localidad de Alcalá de Xivert, provincia de Castellón.\nLa vivienda se distribuye en cuatro plantas. La planta sótano cuenta con garaje para dos vehículos y trastero. La planta baja cuenta con salón-comedor, cocina, aseo una habitación con baño tipo suite y terraza. La planta primera cuenta con tres habitaciones, dos cuartos de baño y una terraza. La planta segunda es una azotea con terraza y barbacoa. Además dispone de jardín delantero y un patio trasero en la planta baja.\nCuenta con puerta de entrada blindada, puertas de paso de madera maciza, armarios empotrados, ventanas correderas Climalit, pintura plástica lisa en techos y paredes, y suelos de plaqueta. Dispone de agua caliente e instalación de aire acondicionado.\nVivienda unifamiliar adosada de tres alturas sobre rasante destinadas a vivienda y una bajo rasante destinada a garaje y trastero del año 2007. Se encuentra ubicada dentro de una urbanización cerrada que dispone de piscina comunitaria, parque infantil y zonas ajardinadas, en el núcleo de Alcossebre, próximo a la Playa Del Moro.\nAlcossebre es un pequeño núcleo costero que consta de diez kilómetros de costa repartidos en cinco playas de gran calidad y diversas calas vírgenes, siendo uno de los pocos pueblos que no se encuentra totalmente urbanizado. Se encuentra en la costa del Azahar, lindando con el término de Peñíscola al norte y el de Torreblanca al sur. Además de contar con la zona costera, se caracteriza por disponer de diversos miradores que ofrecen las montañas pertenecientes al paraje natural de la Sierra de Irta. A lo largo de sus diez kilómetros de costa, destacan cinco grandes playas de arena: la del Cargador, la Romana, la del Moro , la de&apos;Manyetes, y la de Las Fuentes, que poseen la bandera azul como distintivo de calidad. También se pueden encontrar tres calas: Tres playas, la Cala del Moro y la Cala Blanca. Zona con buenas comunicaciones por carrete"
},

*/

var (
	apiURL = "https://api.fotocasa.es/PropertySearch/"
	webURL = "https://www.fotocasa.es"
)

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
		fmt.Printf("error: %s: %s", err, body)
		return l, err
	}
	fmt.Printf("-->> %v\n", l)
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

	err = json.Unmarshal(body, &location)
	if err != nil {
		fmt.Printf("error: %s: %s", err, body)
		return ls, err
	}
	fmt.Printf("-->> %v\n", ls)
	return ls, nil
}

func getHouses(ls LocationSegments, page int) {

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

	body, err := util.Get(searchQuery)
	if err != nil {
		log.Println("error:", err)
	}
	fmt.Println(body)

}
