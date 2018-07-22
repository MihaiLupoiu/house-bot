package telegram

import (
	"log"
)

// Init
func Init(telegramBotID, telegramChatID string) {
	if telegramBotID == "" || telegramChatID == "" {
		log.Println("ERROR: Missing TELEGRAM_BOT_ID or TELEGRAM_CHAT_ID environment variables")
		// os.Exit(1)
	}
}
