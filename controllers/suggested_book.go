package controllers

import (
	"book-organizer/models"

	"github.com/gin-gonic/gin"
)

type SuggestedBookController struct{}

func (s *SuggestedBookController) GetSuggestedBooks(c *gin.Context) {
	currentUser := currentUser(c)
	favoriteCategories := currentUser.FavoriteCategories
	categoryBookModel := new(models.CategoryBookModel)
	suggestedBooks, err := categoryBookModel.GetSuggestedBooks(favoriteCategories)
	if err != nil {
		c.JSON(400, gin.H{"message": "Get suggested books failed"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"suggestedBooks": suggestedBooks})
}
