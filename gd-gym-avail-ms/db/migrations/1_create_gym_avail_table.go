package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	// TODO: deprecate db, compute dynamically in future
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table gym_avail...")
		_, err := db.Exec(`CREATE TABLE gym_avail(
			id SERIAL PRIMARY KEY,
			max_avail INTEGER NOT NULL,
			current_avail INTEGER NOT NULL
		)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table gym_avail...")
		_, err := db.Exec(`DROP TABLE gym_avail`)
		return err
	})
}
