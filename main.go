package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/say", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi",
		})
	})

	router.Run(":5000")
}
