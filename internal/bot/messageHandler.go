package bot

import (
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleUpdate(bot *QuizBot, message *tgBotApi.Message) {
	if message == nil {
		return
	}

	if message.IsCommand() {
		msg := tgBotApi.NewMessage(message.Chat.ID, "")
		switch message.Command() {
		case "start":
			msg.Text = "Привет! Я бот для создания викторин. Напиши /help чтобы узнать мои команды"
		case "help":
			handleHelpMenu(bot, message)
		case "quiz":
			handleQuizMenu(bot, message)
		case "results":
			handleResultsMenu(bot, message)
		case "catalog":
			handleCatalogMenu(bot, message)
		default:
			msg.Text = "I don't know that command"
		}
		bot.api.Send(msg)
	} else {
		if message.Text == "Назад" {
			handleHelpMenu(bot, message)
		}
	}
}

func handleQuizMenu(bot *QuizBot, message *tgBotApi.Message) {
	msg := tgBotApi.NewMessage(message.Chat.ID, "Выбери действие:")

	msg.ReplyMarkup = tgBotApi.NewReplyKeyboard(
		tgBotApi.NewKeyboardButtonRow(
			tgBotApi.NewKeyboardButton("Создать викторину"),
			tgBotApi.NewKeyboardButton("Пройти викторину"),
			tgBotApi.NewKeyboardButton("Мои викторины"),
		),
		tgBotApi.NewKeyboardButtonRow(
			tgBotApi.NewKeyboardButton("Назад"),
		),
	)

	bot.api.Send(msg)
}

func handleHelpMenu(bot *QuizBot, message *tgBotApi.Message) {
	msg := tgBotApi.NewMessage(message.Chat.ID, "Список команд:")

	msg.ReplyMarkup = tgBotApi.NewReplyKeyboard(
		tgBotApi.NewKeyboardButtonRow(
			tgBotApi.NewKeyboardButton("/start"),
			tgBotApi.NewKeyboardButton("/help"),
			tgBotApi.NewKeyboardButton("/quiz"),
			tgBotApi.NewKeyboardButton("/results"),
			tgBotApi.NewKeyboardButton("/catalog"),
		),
	)

	bot.api.Send(msg)
}

func handleResultsMenu(bot *QuizBot, message *tgBotApi.Message) {
	msg := tgBotApi.NewMessage(message.Chat.ID, "Выбери действие:")

	msg.ReplyMarkup = tgBotApi.NewReplyKeyboard(
		tgBotApi.NewKeyboardButtonRow(
			tgBotApi.NewKeyboardButton("Посмотреть все результаты"),
			tgBotApi.NewKeyboardButton("Посмотреть результаты по id викторины"),
			tgBotApi.NewKeyboardButton("Назад"),
		),
	)

	bot.api.Send(msg)
}

func handleCatalogMenu(bot *QuizBot, message *tgBotApi.Message) {
	msg := tgBotApi.NewMessage(message.Chat.ID, "Выбери действие:")

	msg.ReplyMarkup = tgBotApi.NewReplyKeyboard(
		tgBotApi.NewKeyboardButtonRow(
			tgBotApi.NewKeyboardButton("Посмотреть все викторины"),
			tgBotApi.NewKeyboardButton("Посмотреть викторины по тэгу"),
			tgBotApi.NewKeyboardButton("Назад"),
		),
	)

	bot.api.Send(msg)
}
