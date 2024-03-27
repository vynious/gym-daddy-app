package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/ljlimjk10/gym-avail-ms/controllers"
	"github.com/ljlimjk10/gym-avail-ms/db"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env: %s", err)
	}
	dbc := db.InitDBConnection()
	router := gin.Default()
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	gym := router.Group("/api/gym")
	{
		gym.GET("/avail", func(c *gin.Context) {
			controllers.RetrieveCurrentAvail(c, dbc.DB)
		})
		gym.POST("/update-avail", func(c *gin.Context) {
			controllers.UpdateCurrentAvail(c, dbc.DB)
		})
	}

	router.Run("0.0.0.0:3006")
}
