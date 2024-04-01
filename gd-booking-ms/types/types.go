package types

type (
	Localhost        string
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
)

// GormConfig holds the configurations for the gorm database connection

type GormConfig struct {
	Host     Localhost
	Username DatabaseUsername
	Password DatabasePassword
	DBName   DatabaseName
}
