package server

import "github.com/gin-gonic/gin"

func Init() {
	router := NewRouter()
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	router.Run(":3000")
}
