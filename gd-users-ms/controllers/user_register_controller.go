package users_controllers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ljlimjk10/users-ms/auth"
	"github.com/ljlimjk10/users-ms/models/user_models"
	"github.com/ljlimjk10/users-ms/models/user_role_assoc_models"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type RegisterUserPayload struct {
	Username       string `json:"username" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	TelegramHandle string `json:"telegram_handle"`
	RoleID         int    `json:"role_id"`
}

func RegisterUser(c *gin.Context, db *pg.DB) {
	var newUserPayload RegisterUserPayload

	err := c.ShouldBindJSON(&newUserPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPwd, err := auth.HashPassword(newUserPayload.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error. Please contact system admin."})
	}

	currentTime := time.Now()

	newUser := &user_models.User{
		UserID:         uuid.New().String(),
		Username:       newUserPayload.Username,
		Email:          newUserPayload.Email,
		HashedPass:     hashedPwd,
		FirstName:      newUserPayload.FirstName,
		LastName:       newUserPayload.LastName,
		TelegramHandle: newUserPayload.TelegramHandle,
		CreatedAt:      currentTime,
		ModifiedAt:     currentTime,
	}

	// TODO: use tx instead of db in case of user role association creation failure
	if err := user_models.CreateUser(db, newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	newUserRoleAssoc := &user_role_assoc_models.UserRoleAssociation{
		UserID: newUser.UserID,
		RoleID: newUserPayload.RoleID,
	}

	if err := user_role_assoc_models.CreateUserRoleAssoc(db, newUserRoleAssoc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user role association"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}