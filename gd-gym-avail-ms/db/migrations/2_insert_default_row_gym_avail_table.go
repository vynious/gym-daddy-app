package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	// TODO: deprecate db, compute dynamically in future
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("insert default row to table gym_avail...")
		_, err := db.Exec(`INSERT INTO gym_avail (max_avail, current_avail) VALUES (50, 50)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("truncate table gym_avail...")
		_, err := db.Exec(`TRUNCATE TABLE gym_avail`)
		return err
	})
}
