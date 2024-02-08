package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	*tgbotapi.BotAPI
}

var inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Subscribe", "subscribe"),
	),
)

func SpawnBot(uri string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}
	bot.Debug = true
	return &Bot{
		bot,
	}, nil
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {

			callbackData := update.CallbackQuery.Data
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")

			switch callbackData {
			case "subscribe":
				/*
					Implement subscription logic here
					store kv pair {telegramHandle: chatId}
				*/

				//chatId := update.Message.Chat.ID
				//telegramHandle := update.SentFrom().UserName

				msg.Text = "You've subscribed successfully."
				// Optionally, you could acknowledge the callback here with AnswerCallbackQuery
			default:
				msg.Text = "Received unknown command."
			}

			if _, err := b.Send(msg); err != nil {
				fmt.Printf("Error sending message: %v\n", err)
			}
			continue
		}

		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "start":
			msg.Text = "Welcome to gym-daddy."
			msg.ReplyMarkup = inlineKeyboard
		default:
			msg.Text = "I don't know that command."
		}

		if _, err := b.Send(msg); err != nil {
			fmt.Printf("Error sending message: %v\n", err)
		}
	}
}
