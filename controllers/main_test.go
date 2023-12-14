package controllers_test

import (
	"book-organizer/controllers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	health := new(controllers.HealthController)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", health.Status)
	}

	return router
}
