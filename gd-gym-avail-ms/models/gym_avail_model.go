package models

import (
	"errors"

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
	return db.Model(gymAvail).Limit(1).Select()
}

func (gymAvail *GymAvail) DecrementCurrentAvailability(db *pg.DB, qty int) error {
	if gymAvail.CurrentAvail-qty < 0 {
		return errors.New("cannot decrement current availability below 0")
	}

	_, err := db.Model(gymAvail).
		Set("current_avail = current_avail - ?", qty).
		Where("id = ?", gymAvail.ID).
		Update()

	if err == nil {
		gymAvail.CurrentAvail -= qty
	}

	return err
}

func (gymAvail *GymAvail) IncrementCurrentAvailability(db *pg.DB, qty int) error {
	if gymAvail.CurrentAvail+qty < 0 {
		return errors.New("cannot increment current availability above 50")
	}

	_, err := db.Model(gymAvail).
		Set("current_avail = current_avail + ?", qty).
		Where("id = ?", gymAvail.ID).
		Update()

	if err == nil {
		gymAvail.CurrentAvail += qty
	}

	return err
}
