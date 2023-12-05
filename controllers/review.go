package controllers

import (
	"book-organizer/forms"
	"book-organizer/models"

	"github.com/gin-gonic/gin"
)

var reviewModel = new(models.ReviewModel)

type ReviewController struct{}

func (r *ReviewController) CreateReview(c *gin.Context) {
	var data forms.CreateReviewCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()

		return
	}

	user := currentUser(c)
	err := reviewModel.CreateReview(user.ID, data)

	if err != nil {
		c.JSON(400, gin.H{"message": "Error when create review"})
		c.Abort()

		return
	}

	c.JSON(200, gin.H{"message": "Success"})
}

func (r *ReviewController) UpdateReview(c *gin.Context) {
	var data forms.UpdateReviewCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()

		return
	}

	user := currentUser(c)
	err := reviewModel.UpdateReview(user.ID, data)

	if err != nil {
		c.JSON(400, gin.H{"message": "Error when update review"})
		c.Abort()

		return
	}

	c.JSON(200, gin.H{"message": "Success"})
}
