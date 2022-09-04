package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/v1/callback/", func(c *gin.Context) {
		if code, ok := c.GetQuery("code"); ok {
			println(code)
			c.String(200, "callback code ok")
		}
		c.String(400, "callback code failed")
	})
	r.POST("/api/v1/callback/", func(c *gin.Context) {
		code := c.Query("code")
		fmt.Println(code)
		c.String(200, "return code ok")
	})
	r.Run("8088")
}
