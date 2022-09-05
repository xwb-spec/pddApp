package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/v1/callback/", func(c *gin.Context) {
		if code, ok := c.GetQuery("code"); ok {
			Save(ReturnCodeResponse{
				Code:  code,
				State: "xwb",
			})
			c.String(200, "callback code ok")
			return
		}
		c.String(400, "callback code failed")
	})
	r.POST("/api/v1/callback/", func(c *gin.Context) {
		d := GetVal()
		jsonData, err := json.Marshal(&d)
		if err != nil {
			c.JSON(500, "服务器错误")
		}
		c.JSON(200, jsonData)
	})
	r.Run("0.0.0.0:8088")
}
