package main

import (
	"log"
	"os"

	"github.com/MihaiLupoiu/house-bot/lib/logs"
	"github.com/MihaiLupoiu/house-bot/lib/telegram"
	"github.com/coreos/bbolt"
)

func main() {

	logs.Init("[ house-bot ]: ", true)

	telegramBotID := os.Getenv("TELEGRAM_BOT_ID")
	telegramChatID := os.Getenv("TELEGRAM_CHAT_ID")
	telegram.Init(telegramBotID, telegramChatID)

	db, err := bolt.Open("data.db", 0644, bolt.DefaultOptions)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

}
