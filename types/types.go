package types

type (
	MongoURI       string
	DatabaseName   string
	CollectionName string
	ChatID         string
	TelegramHandle string
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
