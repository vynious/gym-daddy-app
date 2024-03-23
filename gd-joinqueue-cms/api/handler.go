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
		Logger: lp,
	}
}

func (qh *QueueHandler) JoinQueue(c *gin.Context) {

	var requestBody struct {
		UserId string `json:"user_Id"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		//go func() {
		//	if err := qh.Logger.SendLog(c, "error", err.Error()); err != nil {
		//		log.Println(err.Error())
		//	}
		//}()
		c.JSON(http.StatusBadRequest, utils.NewError(err))
		return
	}

	userId := requestBody.UserId

	if userId == "" {
		log.Println("Error: no User ID provided")
		//var err error
		//err = errors.New("no user id")
		//go func() {
		//	if err := qh.Logger.SendLog(c, "error", err.Error()); err != nil {
		//		log.Println(err.Error())
		//	}
		//}()
		c.JSON(http.StatusBadRequest, utils.NewError(errors.New("no user id")))
		return
	}

	ticket, err := rpc.GRPCJoinQueue(c, userId)

	//go func() {
	//	if err := qh.Logger.SendLog(c, "default", "created ticket"); err != nil {
	//		log.Println(err.Error())
	//	}
	//}()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, utils.NewError(err))
		//go func() {
		//	if err := qh.Logger.SendLog(c, "error", err.Error()); err != nil {
		//		log.Println(err.Error())
		//	}
		//}()
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ticket})
	return
}

func (qh *QueueHandler) GetUpcomingTicketsInQueue(c *gin.Context) {
	tickets, err := rpc.GRPCGetUpcomingTickets(c)

	//go func() {
	//	if err := qh.Logger.SendLog(c, "default", "get upcoming tickets"); err != nil {
	//		log.Println(err.Error())
	//	}
	//}()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, utils.NewError(err))
		//go func() {
		//	if err := qh.Logger.SendLog(c, "error", err.Error()); err != nil {
		//		log.Println(err.Error())
		//	}
		//}()
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tickets})
	return
}

func (qh *QueueHandler) RetrieveNextInQueue(c *gin.Context) {
	tickets, err := rpc.GRPCGetNextInQueue(c)

	//go func() {
	//	if err := qh.Logger.SendLog(c, "default", "retrieved next in line"); err != nil {
	//		log.Println(err.Error())
	//	}
	//}()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, utils.NewError(err))
		//go func() {
		//	if err := qh.Logger.SendLog(c, "error", err.Error()); err != nil {
		//		log.Println(err.Error())
		//	}
		//}()
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tickets})
	return
}
