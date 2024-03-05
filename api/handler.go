package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vynious/gd-joinqueue-cms/rpc"
	"log"
	"net/http"
)

func JoinQueue(c *gin.Context) {

	var requestBody struct {
		UserId string `json:"user_Id"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": "error"})
		return
	}

	userId := requestBody.UserId

	if userId == "" {
		log.Println("Error: no User ID provided")
		c.JSON(http.StatusBadRequest, gin.H{"data": "error"})
		return
	}

	ticket, err := rpc.GRPCJoinQueue(c, userId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ticket})
	return
}

func GetUpcomingTicketsInQueue(c *gin.Context) {
	tickets, err := rpc.GRPCGetUpcomingTickets(c)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tickets})
	return
}

func RetrieveNextInQueue(c *gin.Context) {
	tickets, err := rpc.GRPCGetNextInQueue(c)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tickets})
	return
}
