package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"user1": "root",
		"user2": "user",
		"user3": "admin",
	}))

	authorized.GET("/secret", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"secret": "root",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
