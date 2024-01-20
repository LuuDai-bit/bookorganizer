package controllers

import (
	"book-organizer/forms"
	"book-organizer/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var bookModel = new(models.BookModel)

type BookController struct{}

func (b *BookController) GetBooks(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Page must be a number"})
		c.Abort()

		return
	}
	search := c.Query("s")

	user := currentUser(c)
	books, err := bookModel.GetBooksByUser(user.ID, page, search)
	if err != nil {
		c.JSON(400, gin.H{"message": "Error when get books"})
		c.Abort()

		return
	}

	total, err := bookModel.GetTotalBookByUser(user.ID, search)
	if err != nil {
		c.JSON(400, gin.H{"message": "Error when get books"})
		c.Abort()

		return
	}

	c.JSON(200, gin.H{"message": "Success", "books": books, "total": total})
}

func (b *BookController) CreateBook(c *gin.Context) {
	var data forms.CreateBookCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()

		return
	}

	user := currentUser(c)
	err := bookModel.CreateBook(user.ID, data)

	if err != nil {
		c.JSON(400, gin.H{"message": "Error when create book"})
		c.Abort()

		return
	}

	c.JSON(200, gin.H{"message": "Success"})
}

func (b *BookController) UpdateBook(c *gin.Context) {
	var data forms.UpdateBookCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()

		return
	}

	user := currentUser(c)
	err := bookModel.UpdateBook(user.ID, data)

	if err != nil {
		c.JSON(400, gin.H{"message": "Error when update book"})
		c.Abort()

		return
	}

	c.JSON(200, gin.H{"message": "Success"})
}

func (b *BookController) Download(c *gin.Context) {
	id := c.Query("id")
	bookId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Download failed"})
		c.Abort()

		return
	}

	user := currentUser(c)

	url, err := bookModel.DownloadEbook(user.ID, bookId)
	if err != nil {
		c.JSON(400, gin.H{"message": "Download failed"})
		c.Abort()

		return
	}

	c.JSON(200, gin.H{"message": "Success", url: url})
}
