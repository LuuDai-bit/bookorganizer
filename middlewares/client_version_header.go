package middlewares

import (
	"github.com/gin-gonic/gin"
)

type ClientVersionHeaderMiddleware struct{}

// TODO: Read this from file. Each time client rebuild dist pump the number up and write it to the file
// For now pump version up manualy
const ClientVersion = "1"

func (a *ClientVersionHeaderMiddleware) AddClientVersionHeader(c *gin.Context) {
	c.Writer.Header().Set("Client-Version", ClientVersion)
}
