package bot

import (
	"fmt"
	"log"
	"os"

	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type QuizBot struct {
	api *tgBotApi.BotAPI
}

func NewQuizBot() (*QuizBot, error) {
	apiKey := os.Getenv("TELEGRAM_BOT_API_TOKEN")
	fmt.Println(apiKey)
	bot, err := tgBotApi.NewBotAPI(apiKey)

	if err != nil {
		return nil, err
	}

	return &QuizBot{
		api: bot,
	}, nil
}

func (bot *QuizBot) Start() {
	log.Printf("Authorized on account %s", bot.api.Self.UserName)

	u := tgBotApi.NewUpdate(0)

	u.Timeout = 60

	updates := bot.api.GetUpdatesChan(u)

	for update := range updates {
		handleUpdate(bot, update.Message)
	}
}
