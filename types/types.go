package types

type (
	MongoURI       string
	DatabaseName   string
	CollectionName string
	ChatID         int64
	TelegramHandle string
	TelegramBotURI string
	EventType      string
	Content        string
)

type UserDocument struct {
	TelHandle TelegramHandle
	ChatId    ChatID
}

type MongoConfig struct {
	Url      MongoURI
	DBName   DatabaseName
	CollName CollectionName
}

type NotificationMessage struct {
	TelegramHandle TelegramHandle
	EventType      EventType
	Content        Content
}
