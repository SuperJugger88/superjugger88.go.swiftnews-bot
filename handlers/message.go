package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"regexp"
	"strings"
	"superjugger88.go.swiftnews-bot/models"
)

func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, user *models.User, inlineKeyboard tgbotapi.InlineKeyboardMarkup) {
	var text string

	switch user.State {
	case models.StateDefault:
		text = "Привет. Как тебя зовут?"
		user.State = models.StateAskName
	case models.StateAskName:
		user.Name = message.Text
		nameToLower := strings.ToLower(user.Name)

		re, _ := regexp.Compile(`на[сз]и`)
		if re.MatchString(nameToLower) {
			text = fmt.Sprintf("Привет %s! Наконец-то ты пришла... Я так рад тебя видеть!\U0001F979", user.Name)
			user.State = models.StateShowKeyboard
		} else {
			text = "Ты не моя мама... Не хочу с тобой разговаривать :("
			user.State = models.StateDefault
		}
	default:
		log.Fatal("Unknown user state")
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	if user.State == models.StateShowKeyboard {
		msg.ReplyMarkup = inlineKeyboard
	}

	if _, err := bot.Send(msg); err != nil {
		log.Fatal(err)
	}
}
