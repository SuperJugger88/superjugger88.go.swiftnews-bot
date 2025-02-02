package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func HandleCallbackQuery(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
	responses := map[string]string{
		"getPhrase": "Вы выбрали фразу дня",
		"orderGift": "Вы выбрали заказ подарка",
	}

	if response, exists := responses[callbackQuery.Data]; exists {
		answer := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, response)
		if _, err := bot.Send(answer); err != nil {
			log.Fatal(err)
		}
	}
}
