package controllers

import (
	"book-organizer/models"
	"book-organizer/supports"
	"time"

	"github.com/gin-gonic/gin"
)

type CalendarDisplayerController struct{}

var calendarModel = new(models.CalendarModel)

// If startDate is not provided set startDate to the first day of current month
// If endDate is not provided set endDate to 30 days from the first day
func (cd *CalendarDisplayerController) GetCalendars(c *gin.Context) {
	timeSupport := new(supports.TimeSupport)
	var startDate, endDate time.Time
	var err error
	if c.Param("startDate") == "" {
		startDate = timeSupport.FirstDayOfMonth(time.Now())
	} else {
		startDate, err = time.Parse(time.DateOnly, c.Param("startDate"))

		if err != nil {
			c.JSON(400, gin.H{"message": "startDate is not a Date"})
			c.Abort()

			return
		}
	}

	if c.Param("endDate") == "" {
		endDate = startDate.AddDate(0, 0, 30)
	} else {
		endDate, err = time.Parse(time.DateOnly, c.Param("endDate"))

		if err != nil {
			c.JSON(400, gin.H{"message": "endDate is not a Date"})
			c.Abort()

			return
		}
	}

	currentUser := currentUser(c)
	calendars, err := calendarModel.GetCalendars(currentUser.ID, startDate, endDate)

	if err != nil {
		c.JSON(400, gin.H{"message": "Error when fetch calendars"})
		c.Abort()

		return
	}

	c.JSON(200, gin.H{"message": "Success", "calendars": calendars})
}
