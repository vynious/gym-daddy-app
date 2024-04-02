package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ljlimjk10/gym-avail-ms/controllers"
	"github.com/ljlimjk10/gym-avail-ms/db"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env: %s", err)
	}
	dbc := db.InitDBConnection()
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	gym := router.Group("/api/gym")
	{
		gym.Use(AuthenticateDefault)
		gym.GET("/avail", func(c *gin.Context) {
			controllers.RetrieveCurrentAvail(c, dbc.DB)
		})
		gym.Use(AuthenticateAdmin)
		gym.POST("/update-avail", func(c *gin.Context) {
			controllers.UpdateCurrentAvail(c, dbc.DB)
		})
	}

	router.Run("0.0.0.0:3006")
}

// Middleware
func AuthenticateDefault(c *gin.Context) {
	log.Println("authenticating user...")

	req, err := http.NewRequestWithContext(c, "GET", "http://user-ms:3005/api/users/validatejwt/default", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "try again later."})
		c.Abort()
		return
	}

	// extract the "Bearer XXXX" and set as header
	token := c.GetHeader("Authorisation")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing auth token please login."})
		c.Abort()
		return
	}

	req.Header.Set("Authorisation", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "try again later."})
		c.Abort()
		return
	}

	defer resp.Body.Close()

	log.Printf("res %+v", resp)
	if resp.Status == http.StatusText(http.StatusUnauthorized) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorised. Go be authorised den come back"})
		c.Abort()
		return
	}
	c.Next()
}

func AuthenticateAdmin(c *gin.Context) {
	log.Println("authenticating admin user...")

	req, err := http.NewRequestWithContext(c, "GET", "http://user-ms:3005/api/users/validatejwt/admin", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "try again later."})
		c.Abort()
		return
	}

	// extract the "Bearer XXXX" and set as header
	token := c.GetHeader("Authorisation")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing auth token please login."})
		c.Abort()
		return
	}

	req.Header.Set("Authorisation", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "try again later."})
		c.Abort()
		return
	}

	if resp.StatusCode == 401 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "only admin-level access"})
		c.Abort()
		return
	}

	defer resp.Body.Close()

	log.Printf("res %+v", resp)
	if resp.Status == http.StatusText(http.StatusUnauthorized) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorised. Go be authorised den come back"})
		c.Abort()
		return
	}
	c.Next()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Authorisation")

		// Add this line to allow preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
