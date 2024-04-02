package types

type (
	MongoURI       string
	DatabaseName   string
	CollectionName string
	
	TelegramBotURI string
	EventType      string
	Content        string
)



type MongoConfig struct {
	Url      MongoURI
	DBName   DatabaseName
	CollName CollectionName
}

type NotificationMessage struct {
	TelegramHandle string
	EventType      EventType
	Content        Content
}
