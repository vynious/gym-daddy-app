package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			fmt.Println("altering username col of users table...")
			_, err := db.Exec(`ALTER TABLE users
			ADD CONSTRAINT users_username_key UNIQUE (username)
		`)
			return err
		},
		func(db migrations.DB) error {
			fmt.Println("removing unique constraint from username col of users table...")
			_, err := db.Exec(`ALTER TABLE users
			DROP CONSTRAINT IF EXISTS users_username_key
		`)
			return err
		},
	)
}
