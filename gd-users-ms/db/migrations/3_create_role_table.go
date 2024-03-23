package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table roles...")
		_, err := db.Exec(`CREATE TABLE roles(
			role_id SERIAL PRIMARY KEY,
			role_name VARCHAR(255) UNIQUE NOT NULL,
			description VARCHAR(255) NOT NULL
		)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table roles...")
		_, err := db.Exec(`DROP TABLE roles`)
		return err
	})
}
