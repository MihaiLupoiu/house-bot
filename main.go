package main

import (
	"log"
	"os"

	"github.com/MihaiLupoiu/house-bot/lib/telegram"
	"github.com/MihaiLupoiu/house-bot/lib/util"
	"github.com/coreos/bbolt"
)

func main() {

	util.InitLog("[ house-bot ]: ", true)

	// TODO: Get from configurtation TELEGRAM_BOT_ID and TELEGRAM_CHAT_ID
	telegramBotID := os.Getenv("TELEGRAM_BOT_ID")
	telegramChatID := os.Getenv("TELEGRAM_CHAT_ID")
	telegram.Init(telegramBotID, telegramChatID)

	//TODO: Get from configuration db name
	db, err := bolt.Open("data.db", 0644, bolt.DefaultOptions)
	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()
}
