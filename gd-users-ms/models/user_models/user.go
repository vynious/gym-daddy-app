package user_models

import (
	"time"
	"log"
	"github.com/go-pg/pg/v10"
)

type User struct {
	tableName  struct{}  `pg:"users"`
	UserID     string       `pg:"user_id, pk"`
	Username   string    `pg:"username, unique, notnull"`
	Email      string    `pg:"email, unique, notnull"`
	HashedPass string    `pg:"hashed_pass, notnull"`
	FirstName  string    `pg:"first_name"`
	LastName   string    `pg:"last_name"`
	TelegramHandle string `pg:"telegram_handle, unique, notnull"`
	CreatedAt  time.Time `pg:"created_at"`
	ModifiedAt time.Time `pg:"modified_at"`
}

func CreateUser(db *pg.DB, newUser *User) error {
	_, err := db.Model(newUser).Insert()
	return err
}

func GetUser(db *pg.DB, username string) (*User, error) {
	user := new(User)
	err := db.Model(user).Where("username = ?", username).Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllUsers(db *pg.DB) (*[]User, error) {
	var allUsers []User
	
	err := db.Model(&allUsers).Select()
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	return &allUsers, nil
}

func GetUserById(db *pg.DB, username string) (*User, error) {
	user := new(User)
	err := db.Model(user).Where("user_id = ?", username).Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}


func GetTelegramHandleByUsername(db *pg.DB, userId string) (string, error) {
	user := new(User)
	err := db.Model(user).Where("user_id = ?", userId).Select()
	if err != nil {
		return "", err
	}
	return user.TelegramHandle, nil
}