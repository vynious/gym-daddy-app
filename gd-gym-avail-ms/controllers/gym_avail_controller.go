package controllers

import (
	"errors"
	"github.com/ljlimjk10/gym-avail-ms/rpc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"

	"github.com/ljlimjk10/gym-avail-ms/models"
)

// type AvailabilityResponse struct {
// 	CurrentAvail
// }

type UpdateCurrentAvailRequestBody struct {
	UpdateType string `json:"update_type"`
	Quantity   int    `json:"quantity"`
}

func RetrieveCurrentAvail(c *gin.Context, db *pg.DB) {
	gymAvail := new(models.GymAvail) // Allocate memory for gymAvail

	err := gymAvail.GetCurrentAvailability(db)
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "availability record not found."})
		} else {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error. Please contact system admin."})
			return
		}
	}

	c.JSON(http.StatusOK, gymAvail.CurrentAvail)
}

func UpdateCurrentAvail(c *gin.Context, db *pg.DB) {
	var reqBody *UpdateCurrentAvailRequestBody

	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gymAvail := new(models.GymAvail)

	
	if reqBody.UpdateType == "increment" {
		if err = gymAvail.IncrementCurrentAvailability(db, reqBody.Quantity); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else if reqBody.UpdateType == "decrement" {
		if err = gymAvail.DecrementCurrentAvailability(db, reqBody.Quantity); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "provide update_type (increment/decrement)"})
		return
	}

	go func() {
		if _, err := rpc.GRPCGetNextInQueue(c); err != nil {
			log.Println(err.Error())
		}
		if _, err := rpc.GRPCGetUpcomingTickets(c); err != nil {
			log.Println(err.Error())
		}
	}()
	
	c.JSON(http.StatusCreated, gymAvail.CurrentAvail)
}
