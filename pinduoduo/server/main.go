package server

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/api/v1/callback", func(c *gin.Context) {
		c.String(200, "callback ok")
	})
	r.Run("8088")
}
