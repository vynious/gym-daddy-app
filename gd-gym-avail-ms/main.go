package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
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
		gym.Use(Authenticate)
		gym.GET("/avail", func(c *gin.Context) {
			controllers.RetrieveCurrentAvail(c, dbc.DB)
		})
		gym.POST("/update-avail", func(c *gin.Context) {
			controllers.UpdateCurrentAvail(c, dbc.DB)
		})
	}

	router.Run("0.0.0.0:3006")
}



// Middleware
func Authenticate(c *gin.Context) {
	log.Println("authenticating user...")

	req, err := http.NewRequestWithContext(c, "GET", "http://user-ms:3005/api/users/validatejwt", nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "try again later.")
		c.Abort()
		return
	}

	// extract the "Bearer XXXX" and set as header
	token := c.GetHeader("Authorisation")
	if token == "" {
		c.String(http.StatusBadRequest, "missing auth token please login.")
		c.Abort()
		return
	}

	req.Header.Set("Authorisation", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "try again later.")
		c.Abort()
		return
	}

	defer resp.Body.Close()

	log.Printf("res %+v", resp)
	if resp.Status == http.StatusText(http.StatusUnauthorized) {
		c.String(http.StatusUnauthorized, "Unauthorised. Go be authorised den come back")
		c.Abort()
		return
	}
	c.Next()
}
