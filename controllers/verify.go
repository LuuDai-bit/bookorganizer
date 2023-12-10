package controllers

import (
	"book-organizer/forms"
	"book-organizer/models"

	"github.com/gin-gonic/gin"
)

var verifyModel = new(models.VerifyModel)

type VerifyController struct{}

func (v *VerifyController) SendVerifyCode(c *gin.Context) {
	var data forms.SendVerifyCodeCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()

		return
	}

	err := verifyModel.SendVerifyCode(data.Email)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(201, gin.H{"message": "Please check your email for otp code"})
}
