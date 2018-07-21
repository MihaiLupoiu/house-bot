package main

import (
	"log"
	"os"
)

func main() {
	telegramBotID := os.Getenv("TELEGRAM_BOT_ID")
	telegramChatID := os.Getenv("TELEGRAM_CHAT_ID")

	if telegramBotID == "" || telegramChatID == "" {
		log.Println("ERROR: Missing TELEGRAM_BOT_ID or TELEGRAM_CHAT_ID environment variables")
		os.Exit(1)
	}
}
