package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vynious/gd-joinqueue-cms/logger"
	"github.com/vynious/gd-joinqueue-cms/rpc"
	"github.com/vynious/gd-joinqueue-cms/utils"
	"log"
	"net/http"
)

type QueueHandler struct {
	Logger *logger.LogProducer
}

func SpawnQueueHandler(lp *logger.LogProducer) *QueueHandler {
	return &QueueHandler{
		Logger: lp, // not in use...
	}
}

func (qh *QueueHandler) JoinQueue(c *gin.Context) {

	var requestBody struct {
		UserID string `json:"user_id"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewError(err))
		return
	}

	userId := requestBody.UserID

	if userId == "" {
		log.Println("Error: no User ID provided")
		c.JSON(http.StatusBadRequest, utils.NewError(errors.New("no user id")))
		return
	}

	// join queue 
	ticket, err := rpc.GRPCJoinQueue(c, userId)

	// send notification 
	go rpc.GRPCSendNotification(c, nil, ticket, "Join-Queue")

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, utils.NewError(err))

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ticket})
}

func (qh *QueueHandler) GetUpcomingTicketsInQueue(c *gin.Context) {
	tickets, err := rpc.GRPCGetUpcomingTickets(c)


	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, utils.NewError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tickets})
}

func (qh *QueueHandler) RetrieveNextInQueue(c *gin.Context) {
	tickets, err := rpc.GRPCGetNextInQueue(c)


	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, utils.NewError(err))

		return
	}


	c.JSON(http.StatusOK, gin.H{"data": tickets})
}
