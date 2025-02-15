package main

import (
	"KandeHarsha/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	baseGroup := r.Group("api/v1")

	authRouter := baseGroup.Group("/auth")
	handler.NewAuthHandler(authRouter)

	r.Run("0.0.0.0:4000") // listen and serve on port 4000
}
