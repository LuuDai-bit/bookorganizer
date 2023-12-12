package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"book-organizer/controllers"
	"book-organizer/models"
)

func main() {
	router := gin.Default()

	health := new(controllers.HealthController)
	user := new(controllers.UserController)
	session := new(controllers.SessionController)
	book := new(controllers.BookController)
	review := new(controllers.ReviewController)
	statistic := new(controllers.StatisticController)
	verify := new(controllers.VerifyController)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", health.Status)
		v1.POST("/signup", user.Signup)
		v1.POST("/signin", session.SignIn)
		v1.DELETE("/logout", session.LogOut)
		v1.PATCH("/users/change_password", user.UpdatePassword)
		v1.POST("/verify/send", verify.SendVerifyCode)
	}
	v1.Use(AuthRequired)
	{
		v1.GET("/users/me", user.ShowDetail)
		v1.GET("/books/:page", book.GetBooks)
		v1.POST("/books/create", book.CreateBook)
		v1.PATCH("/books/update", book.UpdateBook)
		v1.POST("/reviews/create", review.CreateReview)
		v1.PATCH("/reviews/update", review.UpdateReview)
		v1.GET("/statistic/books/read/:year", statistic.CountBook)
		v1.GET("/statistic/categories/favorite", statistic.GetFavoriteCateogries)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	router.Run(":3000")
}

func AuthRequired(c *gin.Context) {
	token := c.GetHeader("Token")
	sessionModel := new(models.SessionModel)
	_, err := sessionModel.FindOne(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	c.Next()
}
