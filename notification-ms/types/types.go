package types

type (
	Localhost        string
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
	KafkaURL         string
	KafkaTopic       string
)

type Message struct {
	userId  string
	event   string // type of message to send to the user (event == topic)
	content string
}

type GormConfig struct {
	Host     Localhost
	Username DatabaseUsername
	Password DatabasePassword
	DBName   DatabaseName
}

type KafkaWriterConfig struct {
	Url   KafkaURL
	Topic KafkaTopic
}
