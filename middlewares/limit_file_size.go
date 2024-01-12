package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const MaxBodyBytes = 1024 * 1024 * 5 // 5MB

type LimitFileSizeMiddleware struct{}

func (lfs *LimitFileSizeMiddleware) BodySizeMiddleware(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	c.Request.Body = http.MaxBytesReader(w, c.Request.Body, MaxBodyBytes)

	c.Next()
}
