package main

import (
	"context"
	"flag"
	"log"

	"github.com/MihaiLupoiu/house-bot/lib/telegram"
	"github.com/MihaiLupoiu/house-bot/lib/util"
	"github.com/MihaiLupoiu/house-bot/models"
	"github.com/MihaiLupoiu/house-bot/providers"

	"github.com/coreos/bbolt"
)

func main() {

	// GET configuration
	configFilePath := flag.String("configFile", "./config.json", "JSON config file to read.")
	flag.Parse()
	config := util.GetConfigurationFile(*configFilePath)

	util.InitLog("[ house-bot ]: ", config.Debug)

	// initialize Database
	db, err := bolt.Open(config.Database, 0644, bolt.DefaultOptions)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	// Create acommunication Channel
	houseChan := make(chan *models.House, 1000)
	ctx, ctxCancel := context.WithCancel(context.Background())

	fotocasa.Init(db, houseChan, config.Fotocasa, config.Filters)
	go fotocasa.TickerCheck(ctx, ctxCancel)

	telegram.Init(config.Telegram, houseChan)
	telegram.RunBot(ctx, ctxCancel)

}
