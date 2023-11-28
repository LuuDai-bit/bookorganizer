package controllers

import (
	"book-organizer/forms"
	"book-organizer/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var bookModel = new(models.BookModel)

type BookController struct{}

func (b *BookController) GetBooks(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Page must be a number"})
	}

	user := current_user(c)
	books := bookModel.GetBooksByUser(user.ID, page)

	c.JSON(200, gin.H{"message": "Success", "books": books})
}

func (b *BookController) CreateBook(c *gin.Context) {
	var data forms.CreateBookCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()

		return
	}

	user := current_user(c)
	err := bookModel.CreateBook(user.ID, data)

	if err != nil {
		c.JSON(400, gin.H{"message": "Error when create book"})
	}

	c.JSON(200, gin.H{"message": "Success"})
}
