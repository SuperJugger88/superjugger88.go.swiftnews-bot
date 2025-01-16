package main

import (
	"fmt"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
	"os"
	"superjugger88.go.swiftnews-bot/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	botToken := config.TELEGRAM_APITOKEN

	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel
	// (more on configuration in examples/updates_long_polling/main.go)
	updates, _ := bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot, updates)

	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		chatID := tu.ID(message.Chat.ID)
		_, _ = bot.CopyMessage(
			tu.CopyMessage(chatID, chatID, message.MessageID),
		)
	})

	defer bh.Stop()

	// Stop reviving updates from update channel
	defer bot.StopLongPolling()

	bh.Start()
}
