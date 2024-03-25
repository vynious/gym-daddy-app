package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

// func init() {
// 	migrations.MustRegisterTx(func(db migrations.DB) error {
// 		fmt.Println("creating table users...")
// 		_, err := db.Exec(`CREATE TABLE users(
// 			user_id SERIAL PRIMARY KEY,
// 			username VARCHAR(255) NOT NULL,
// 			email VARCHAR(255) UNIQUE NOT NULL,
// 			hashed_pass VARCHAR(255) NOT NULL,
// 			first_name VARCHAR(255),
// 			last_name VARCHAR(255),
// 			telegram_handle VARCHAR(255),
// 			created_at TIMESTAMP NOT NULL,
// 			modified_at TIMESTAMP
// 		)`)
// 		return err
// 	}, func(db migrations.DB) error {
// 		fmt.Println("dropping table users...")
// 		_, err := db.Exec(`DROP TABLE users`)
// 		return err
// 	})
// }

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table users...")
		_, err := db.Exec(`CREATE TABLE users(
			user_id VARCHAR(255) PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			hashed_pass VARCHAR(255) NOT NULL,
			first_name VARCHAR(255),
			last_name VARCHAR(255),
			telegram_handle VARCHAR(255) UNIQUE,
			created_at TIMESTAMP NOT NULL,
			modified_at TIMESTAMP
		)`)
		if err != nil {
			return err
		}

		fmt.Println("inserting dummy data into users table...")
		_, err = db.Exec(`INSERT INTO users (user_id, username, email, hashed_pass, first_name, last_name, telegram_handle, created_at, modified_at) VALUES ('shawn-thiah','shawn-thiah', 'shawn@example.com', 'hashedpassword', 'Shawn', 'Thiah', 'shawntyw', NOW(), NOW())`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table users...")
		_, err := db.Exec(`DROP TABLE users`)
		return err
	})
}
