package main

import (
	"log"

	bot "github.com/alexey0101/quiz_simple_bot/internal/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	bot, err := bot.NewQuizBot()

	if err != nil {
		log.Fatalf("Error creating bot: %s", err)
	}

	bot.Start()
}
