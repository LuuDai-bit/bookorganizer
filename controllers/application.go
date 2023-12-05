package controllers

import (
	"github.com/gin-gonic/gin"

	"book-organizer/models"
)

func currentUser(c *gin.Context) models.User {
	token := c.GetHeader("Token")
	session, _ := sessionModel.FindOne(token)
	user, err := userModel.FindOne(session.UserID)

	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "User not found"})
		return user
	}

	return user
}
