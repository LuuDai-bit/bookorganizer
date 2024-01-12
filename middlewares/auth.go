package middlewares

import (
	"book-organizer/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthMiddleware struct{}

func (a *AuthMiddleware) AuthRequired(c *gin.Context) {
	sessionModel := new(models.SessionModel)
	userModel := new(models.UserModel)

	token := c.GetHeader("Token")
	session, _ := sessionModel.FindOne(token)
	_, err := userModel.FindOne(session.UserID)

	currentTime := primitive.NewDateTimeFromTime(time.Now())
	if err != nil || session.ExpireTime < currentTime {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
	}
}
