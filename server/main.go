package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/blade-runner", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "I've seen things you people wouldn't believe...",
			"movie":   "Blade Runner",
			"year":    1982,
		})
	})

	r.Run(":8080")
}