package controllers

import (
	"book-organizer/forms"
	"book-organizer/models"

	"github.com/gin-gonic/gin"
)

var userModel = new(models.UserModel)

type UserController struct{}

func (u *UserController) Signup(c *gin.Context) {
	var data forms.SignupUserCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()

		return
	}

	err := userModel.Signup(data)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(201, gin.H{"message": "New user account registered"})
}

func (u *UserController) UpdatePassword(c *gin.Context) {
	var data forms.UpdateUserPasswordCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()

		return
	}

	err := userModel.UpdatePassword(data)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Password changed"})
}

func (u *UserController) ShowDetail(c *gin.Context) {
	user := current_user(c)

	c.JSON(200, gin.H{"message": "Success", "user": user})
}
