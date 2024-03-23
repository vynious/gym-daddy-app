package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/ljlimjk10/gym-avail-ms/controllers"
	"github.com/ljlimjk10/gym-avail-ms/db"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env: %s", err)
	}
	db := db.InitDBConnection()
	router := gin.Default()
	gym := router.Group("/gym")
	{
		gym.GET("/avail", func(c *gin.Context) {
			controllers.RetrieveCurrentAvail(c, db.DB)
		})
		gym.POST("/update-avail", func(c *gin.Context) {
			controllers.UpdateCurrentAvail(c, db.DB)
		})
	}

	router.Run("0.0.0.0:3006")

}
