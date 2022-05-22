package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
			"time":    time.Now().Format("2006-01-02 15:04:05"),
			"status":  200,
		})
	})
	_ = r.Run()
}
