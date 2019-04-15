package apps

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// http://localhost:8080
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world")
	})
	// http://localhost:8080/hoge
	router.GET("/hoge", func(c *gin.Context) {
		c.String(200, "fuga")
	})
	router.Run(":8080")
}