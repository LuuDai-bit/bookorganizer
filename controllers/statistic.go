package controllers

import (
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

	statistic := new(services.BookStatistic)
	result := statistic.NumberOfBookRead(year)

	c.JSON(200, gin.H{"message": "Success", "result": result})
}

func (s *StatisticController) GetFavoriteCateogries(c *gin.Context) {
	statistic := new(services.BookStatistic)
	result := statistic.FavoriteCategories()

	c.JSON(200, gin.H{"message": "Success", "result": result})
}
