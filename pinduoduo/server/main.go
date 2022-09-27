package server

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/v1/callback/", func(c *gin.Context) {
		code, err := c.GetQuery("code")
		if err != nil {
			c.String(400, "callback code failed")
			return
		}
		state, err := c.GetQuery("state")
		if err != nil {
			c.String(400, "callback code failed")
			return
		}
		Save(ReturnCodeResponse{
			Code:  code,
			State: state,
		})
		c.String(200, "callback code ok")
		return
	})
	r.POST("/api/v1/callback/", func(c *gin.Context) {
		d := GetVal()
		c.JSON(200, d)
	})
	r.Run("0.0.0.0:8088")
}
