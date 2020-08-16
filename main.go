package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	go h.run()

	gin.SetMode(gin.DebugMode)

	var router = gin.New()
	router.LoadHTMLFiles("./views/index.html")

	router.GET("/room/:roomId/:userName", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId/:userName", func(c *gin.Context) {
		roomId := c.Param("roomId")
		userName := c.Param("userName")
		serveWs(c.Writer, c.Request, roomId, userName)
	})

	router.Run("0.0.0.0:3001")
}
