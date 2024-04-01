package models

import (
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
)

type GymAvail struct {
	tableName    struct{} `pg:"gym_avail"`
	ID           int      `pg:"id, pk"`
	MaxAvail     int      `pg:"max_avail, notnull"`
	CurrentAvail int      `pg:"current_avail, notnull"`
}

// TODO: db should be deprecated. compute availability dynamically

func (gymAvail *GymAvail) GetCurrentAvailability(db *pg.DB) error {
	err := db.Model(gymAvail).Limit(1).Select()
	if err != nil {
		fmt.Printf("Error fetching current availability: %v", err)
		return err
	}

	fmt.Printf("Fetched gym availability: %+v", gymAvail)
	return nil
}

func (gymAvail *GymAvail) DecrementCurrentAvailability(db *pg.DB, qty int) error {


	gymAvail.ID = 1

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()


	err = tx.Model(gymAvail).WherePK().Select()
	if err != nil {
		return err
	}

	if gymAvail.CurrentAvail-qty < 0 {
		return errors.New("cannot decrement current availability below 0")
	}

	_, err = tx.Model(gymAvail).
		Set("current_avail = current_avail - ?", qty).
		WherePK().
		Update()
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	gymAvail.CurrentAvail -= qty

	return nil
}

func (gymAvail *GymAvail) IncrementCurrentAvailability(db *pg.DB, qty int) error {
	gymAvail.ID = 1


	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()


	err = tx.Model(gymAvail).WherePK().Select()
	if err != nil {
		return err
	}

	if gymAvail.CurrentAvail+qty > 50 {
		return errors.New("cannot increment current availability above 50")
	}

	_, err = tx.Model(gymAvail).
		Set("current_avail = current_avail + ?", qty).
		WherePK().
		Update()
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	gymAvail.CurrentAvail += qty

	return nil
}