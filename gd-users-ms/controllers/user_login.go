package users_controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/ljlimjk10/users-ms/auth"
	"github.com/ljlimjk10/users-ms/models/user_models"
	"github.com/ljlimjk10/users-ms/models/user_role_assoc_models"
	"github.com/ljlimjk10/users-ms/types"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func LoginUser(c *gin.Context, db *pg.DB, jwtAuthService *auth.JwtService) {
	var loginPayload types.LoginPayload
	err := c.ShouldBindJSON(&loginPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Logging in...")

	targetUser, err := user_models.GetUser(db, loginPayload.Username)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user record found. Please register or try again with another set of credentials"})
		return
	}
	if !auth.ComparePasswordHash(loginPayload.Password, targetUser.HashedPass) {
		fmt.Println("Invalid password.")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password. Please try again."})
		return
	}

	user, err := user_models.GetUser(db, loginPayload.Username)
	if err != nil {
		log.Println(err)
		if errors.Is(err, pg.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found."})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error. Please contact system admin."})
			return
		}
	}

	roleId, err := user_role_assoc_models.GetUserRoleID(db, user.UserID)

	signedToken, err := jwtAuthService.TokenGenerate(loginPayload.Username, roleId)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error. Please contact system admin."})
		return
	}

	c.JSON(http.StatusOK, signedToken)
}
