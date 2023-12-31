package controllers

import (
	"book-organizer/models"

	"book-organizer/forms"

	"github.com/gin-gonic/gin"
)

var sessionModel = new(models.SessionModel)

type SessionController struct{}

func (s *SessionController) SignIn(c *gin.Context) {
	var data forms.SignInUserCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()

		return
	}

	token, err := userModel.SignIn(data)

	if err != nil {
		needVerify := err.Error() == "Email not verify"

		c.JSON(400, gin.H{"message": err.Error(), "needVerify": needVerify})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Logged in", "token": token})
}

func (s *SessionController) LogOut(c *gin.Context) {
	token := c.GetHeader("Token")

	result, err := sessionModel.Destroy(token)

	if err != nil || result.DeletedCount < 1 {
		c.JSON(400, gin.H{"message": "Error"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Logged out"})
}
