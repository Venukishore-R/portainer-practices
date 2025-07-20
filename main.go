package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Ok...",
		})
	})

	r.Run(":5000")
}
