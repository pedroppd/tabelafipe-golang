package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "API is health",
		})
	})

	server.Run(":8000")
}
