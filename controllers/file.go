package controllers

import (
	"book-organizer/models"
	"book-organizer/services"

	"github.com/gin-gonic/gin"
)

type FileController struct{}

func (f *FileController) UploadSingleFile(c *gin.Context) {
	file, _ := c.FormFile("file")

	s3Handler := new(services.S3Handler)
	key, err := s3Handler.UploadFile(file)
	if err != nil {
		c.JSON(400, gin.H{"message": "Upload failed"})
		c.Abort()

		return
	}

	fileModel := new(models.FileModel)
	err = fileModel.CreateFile(key)
	if err != nil {
		c.JSON(400, gin.H{"message": "Upload failed"})
		c.Abort()

		return
	}

	c.JSON(200, gin.H{"message": "Success", "key": key})
}
