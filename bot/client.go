package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	*tgbotapi.BotAPI
}

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
