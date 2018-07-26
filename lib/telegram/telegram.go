package telegram

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/MihaiLupoiu/house-bot/models"
	"github.com/leekchan/accounting"
	"gopkg.in/telegram-bot-api.v4"
)

var (
	BotID     string
	ChannelID int64
	HouseChan chan *models.House
	AC        = accounting.Accounting{Symbol: "€", Precision: 0}
)

// Init
func Init(config models.Telegram, communicationChannel chan *models.House) {
	if config.BotID == "" || config.ChannelID == "" {
		log.Println("ERROR: Missing TELEGRAM_BOT_ID or TELEGRAM_CHAT_ID environment variables")
		os.Exit(1)
	}
	BotID = config.BotID
	ChannelID, _ = strconv.ParseInt(config.ChannelID, 10, 64)
	HouseChan = communicationChannel
}

func RunBot(ctx context.Context, ctxCancel context.CancelFunc) {
	bot, err := tgbotapi.NewBotAPI(BotID)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	for {
		select {
		case house := <-HouseChan:
			txt := fmt.Sprintf("%s\n%s\n%s", house.Title, AC.FormatMoneyInt(house.Price), house.URL)
			fmt.Printf("%s\n", txt)

			//bot.Send(tgbotapi.NewMessage(ChannelID, txt))

		case <-ctx.Done():
			fmt.Println("bye")
			return
		}
	}

	ctxCancel() //unreachable

	/* // TODO: Accept commands
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
	*/
}
