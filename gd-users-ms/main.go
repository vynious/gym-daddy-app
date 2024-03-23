package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/ljlimjk10/users-ms/auth"
	users_controllers "github.com/ljlimjk10/users-ms/controllers"
	"github.com/ljlimjk10/users-ms/db"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env: %s", err)
	}
	db := db.InitDBConnection()
	jwtAuthService := auth.JWTAuthService()
	router := gin.Default()
	users := router.Group("/users")
	{
		users.POST("/register", func(c *gin.Context) {
			users_controllers.RegisterUser(c, db.DB)
		})
		users.POST("/login", func(c *gin.Context) {
			users_controllers.LoginUser(c, db.DB, jwtAuthService)
		})
		users.GET("/allusers", jwtAuthService.JwtAuthMiddleware(), func(c *gin.Context) {
			users_controllers.GetAllUsers(c, db.DB)
		})
		users.GET("/user", jwtAuthService.JwtAuthMiddleware(), func(c *gin.Context) {
			users_controllers.GetUser(c, db.DB)
		})
		users.GET("/validatejwt", jwtAuthService.JwtAuthMiddleware())
	}

	router.Run("0.0.0.0:3005")

}
