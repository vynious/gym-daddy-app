package users_controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/ljlimjk10/users-ms/models/role_models"
	"github.com/ljlimjk10/users-ms/models/user_models"
	"github.com/ljlimjk10/users-ms/models/user_role_assoc_models"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type SanitisedUser struct {
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

func GetUser(c *gin.Context, db *pg.DB) {
	username := c.Query("username")
	user, err := user_models.GetUser(db, username)
	if err != nil {
		log.Println(err)
		if errors.Is(err, pg.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found."})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error. Please contact system admin."})
			return
		}
	}

	userRoleID, err := user_role_assoc_models.GetUserRoleID(db, user.UserID)
	if err != nil {
		log.Println(err)
		if errors.Is(err, pg.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User role not found."})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error. Please contact system admin."})
			return
		}
		return
	}
	sanitisedUser := &SanitisedUser{
		UserID:    user.UserID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      role_models.RoleIDToName(userRoleID),
	}

	c.JSON(http.StatusOK, sanitisedUser)

}

func GetAllUsers(c *gin.Context, db *pg.DB) {
	allUsers, err := user_models.GetAllUsers(db)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error. Please contact system admin."})
		return
	}

	var sanitisedUsers []*SanitisedUser
	for _, user := range allUsers {
		userRoleID, err := user_role_assoc_models.GetUserRoleID(db, user.UserID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error. Please contact system admin."})
			return
		}
		sanitisedUsers = append(sanitisedUsers, &SanitisedUser{
			UserID:    user.UserID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      role_models.RoleIDToName(userRoleID),
		})
	}
	c.JSON(http.StatusOK, sanitisedUsers)
}
