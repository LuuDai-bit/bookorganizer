package main

import (
	"book-organizer/controllers"
	"book-organizer/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	router := setupRouter()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	router.Run(":3000")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	health := new(controllers.HealthController)
	user := new(controllers.UserController)
	session := new(controllers.SessionController)
	book := new(controllers.BookController)
	review := new(controllers.ReviewController)
	statistic := new(controllers.StatisticController)
	verify := new(controllers.VerifyController)
	file := new(controllers.FileController)
	router.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"POST", "PUT", "PATCH", "GET", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		}),
		new(middlewares.LimitFileSizeMiddleware).BodySizeMiddleware,
		gin.Recovery(),
	)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", health.Status)
		v1.POST("/signup", user.Signup)
		v1.POST("/signin", session.SignIn)
		v1.DELETE("/logout", session.LogOut)
		v1.POST("/verify/send", verify.SendVerifyCode)
		v1.POST("/verify/activate", verify.VerifyAccount)
	}
	v1.Use(new(middlewares.AuthMiddleware).AuthRequired)
	{
		v1.GET("/users/me", user.ShowDetail)
		v1.PATCH("/users/change_password", user.UpdatePassword)
		v1.PATCH("/users/change_avatar", user.UpdateAvatar)
		v1.GET("/books/:page", book.GetBooks)
		v1.POST("/books/create", book.CreateBook)
		v1.PATCH("/books/update", book.UpdateBook)
		v1.POST("/reviews/create", review.CreateReview)
		v1.PATCH("/reviews/update", review.UpdateReview)
		v1.GET("/statistic/books/read/:year", statistic.CountBook)
		v1.GET("/statistic/categories/favorite", statistic.GetFavoriteCateogries)
		v1.POST("/file/single", file.UploadSingleFile)
	}

	return router
}
