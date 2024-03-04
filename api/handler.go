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
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}

	userId := requestBody.UserId

	if userId == "" {
		log.Println("Error: no User ID provided")
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}

	if err := rpc.RPCJoinQueue(c, userId); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ""})
	return
}
