package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/ljlimjk10/users-ms/auth"

	users_controllers "github.com/ljlimjk10/users-ms/controllers"
	"github.com/ljlimjk10/users-ms/db"
)

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

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env: %s", err)
	}
	db := db.InitDBConnection()
	jwtAuthService := auth.JWTAuthService()
	router := gin.Default()

	router.Use(CORSMiddleware())

	users := router.Group("/api/users")
	{
		users.POST("/register", func(c *gin.Context) {
			users_controllers.RegisterUser(c, db.DB)
		})
		users.POST("/login", func(c *gin.Context) {
			users_controllers.LoginUser(c, db.DB, jwtAuthService)
		})
		users.GET("/allusers", jwtAuthService.JwtAuthMiddlewareAdmin(), func(c *gin.Context) {
			users_controllers.GetAllUsers(c, db.DB)
		})
		users.GET("/user", jwtAuthService.JwtAuthMiddlewareDefault(), func(c *gin.Context) {
			users_controllers.GetUser(c, db.DB)
		})
		users.GET("/telegram", func(c *gin.Context) {
			users_controllers.GetTelegramHandle(c, db.DB)
		})
		users.GET("/validatejwt/admin", jwtAuthService.JwtAuthMiddlewareAdmin())
		users.GET("/validatejwt/default", jwtAuthService.JwtAuthMiddlewareDefault())
	}
	
	router.Run("0.0.0.0:3005")

}
