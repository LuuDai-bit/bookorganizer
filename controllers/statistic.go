package controllers

import (
	"book-organizer/models"
	"book-organizer/services"

	"github.com/gin-gonic/gin"
)

type StatisticController struct{}

func (s *StatisticController) CountBook(c *gin.Context) {
	year := c.Param("year")

	if year == "" {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()

		return
	}

	user := currentUser(c)
	statistic := new(services.BookStatistic)
	result, err := statistic.NumberOfBookRead(user.ID, year)

	if err != nil {
		c.JSON(400, gin.H{"message": "Failed"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Success", "result": result})
}

func (s *StatisticController) GetFavoriteCateogries(c *gin.Context) {
	statistic := new(services.BookStatistic)
	result, err := statistic.FavoriteCategories()
	if err != nil {
		c.JSON(400, gin.H{"message": "Failed"})
		c.Abort()
		return
	}

	userModel = new(models.UserModel)
	currentUser := currentUser(c)
	err = userModel.UpdateUserFavoriteCategories(result, currentUser.ID)

	if err != nil {
		c.JSON(400, gin.H{"message": "Failed"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Success", "result": result})
}
