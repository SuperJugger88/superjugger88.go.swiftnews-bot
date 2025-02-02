package main

import (
	"log"
	"superjugger88.go.swiftnews-bot/handlers"
	"superjugger88.go.swiftnews-bot/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"superjugger88.go.swiftnews-bot/util"
)

var inlineKeyboard = tgbotapi.InlineKeyboardMarkup{
	InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("Фраза дня", "getPhrase"),
			tgbotapi.NewInlineKeyboardButtonData("Заказать подарочек", "orderGift"),
		},
	},
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	botToken := config.TELEGRAM_APITOKEN

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	var user models.User

	for update := range updates {
		if update.CallbackQuery != nil {
			handlers.HandleCallbackQuery(bot, update.CallbackQuery)
		} else if update.Message != nil {
			handlers.HandleMessage(bot, update.Message, &user, inlineKeyboard)
		}
	}
}
