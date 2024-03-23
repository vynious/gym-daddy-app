package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"

	"github.com/ljlimjk10/gym-avail-ms/models"
	"github.com/ljlimjk10/gym-avail-ms/rpc"
)

// type AvailabilityResponse struct {
// 	CurrentAvail
// }

type UpdateCurrentAvailRequestBody struct {
	UpdateType string `json:"update_type"`
	Quantity   int    `json:"quantity"`
}

func RetrieveCurrentAvail(c *gin.Context, db *pg.DB) {
	var gymAvail *models.GymAvail

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

	var gymAvail *models.GymAvail

	if reqBody.UpdateType == "increment" {
		gymAvail.IncrementCurrentAvailability(db, reqBody.Quantity)
		rpc.GRPCGetNextInQueue(c)
	} else if reqBody.UpdateType == "decrement" {
		gymAvail.DecrementCurrentAvailability(db, reqBody.Quantity)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "provide update_type (increment/decrement)"})
	}

}
