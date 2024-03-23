package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("seeding roles...")
		_, err := db.Exec(`INSERT INTO roles VALUES 
			(1, 'user', 'gym user'),
			(2, 'admin', 'gym administrator')
			`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("truncating roles...")
		_, err := db.Exec(`TRUNCATE roles`)
		return err
	})
}
