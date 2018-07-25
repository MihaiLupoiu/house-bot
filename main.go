package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/MihaiLupoiu/house-bot/lib/util"
	"github.com/MihaiLupoiu/house-bot/providers"

	"github.com/coreos/bbolt"
)

func main() {

	util.InitLog("[ house-bot ]: ", true)

	// GET conffigration
	configFilePath := flag.String("configFile", "./config/config.json", "JSON config file to read.")
	flag.Parse()

	config := util.GetConfigurationFile(*configFilePath)

	fmt.Println(config)

	// telegram.Init(config.Telegram.BotID, config.Telegram.ChannelID)

	db, err := bolt.Open(config.Database, 0644, bolt.DefaultOptions)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	//	location, _ := fotocasa.GetCityLocation("castellon")
	//	locDet, _ := fotocasa.GetCombinedLocationIds(location)
	//	fotocasa.GetHouses(locDet, 1)

	fotocasa.TickerCheck(db)

}
