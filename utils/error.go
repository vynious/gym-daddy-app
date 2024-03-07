package utils

import "github.com/gin-gonic/gin"

func NewError(err error) gin.H {
	return gin.H{"message": err.Error()}
}
