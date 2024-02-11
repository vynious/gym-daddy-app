package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vynious/gd-telemessenger-ms/db"
	"github.com/vynious/gd-telemessenger-ms/types"
	"log"
	"os"
)

type Bot struct {
	*tgbotapi.BotAPI
	Database *db.Repository
}

var inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Subscribe", "subscribe"),
	),
)

func LoadBotConfig() types.TelegramBotURI {
	telegramUri := os.Getenv("TELEGRAM_BOT_URI")
	if telegramUri == "" {
		log.Fatalf("missing uri for telegram bot")
	}
	return types.TelegramBotURI(telegramUri)
}

func SpawnBot(cfg types.TelegramBotURI, repo *db.Repository) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(string(cfg))
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}
	bot.Debug = true
	return &Bot{
		bot,
		repo,
	}, nil
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			callbackData := update.CallbackQuery.Data
			chatId := update.CallbackQuery.Message.Chat.ID
			msg := tgbotapi.NewMessage(chatId, "")
			telegramHandle := update.CallbackQuery.From.UserName

			switch callbackData {
			case "subscribe":
				if err := b.Database.CreateSubscription(types.TelegramHandle(telegramHandle), types.ChatID(chatId)); err != nil {
					log.Println("Error creating subscription:", err)
					msg.Text = "Oh no! Something went wrong with the server, try again."
				} else {
					msg.Text = "You've subscribed successfully."
				}

				// Send response message
				if _, err := b.Send(msg); err != nil {
					log.Println("Failed to send message:", err)
				}
			default:
				msg.Text = "Received unknown command."
				if _, err := b.Send(msg); err != nil {
					log.Println("Failed to send message:", err)
				}
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
			log.Println("Error sending message:", err)
		}
	}
}

func (b *Bot) SendNotification(chatId types.ChatID, message *types.NotificationMessage) error {
	msg := tgbotapi.NewMessage(int64(chatId), string(message.Content))
	_, err := b.Send(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}
