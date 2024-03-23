package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table user role association...")
		_, err := db.Exec(`CREATE TABLE user_role_association(
			user_role_id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			role_id INT NOT NULL,
			CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
			CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES roles(role_id) ON DELETE CASCADE
		)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table roles...")
		_, err := db.Exec(`DROP TABLE roles`)
		return err
	})
}
