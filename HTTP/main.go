package main

import (
	// "http/service"

	"github.com/gin-gonic/gin"
)

// var (
// 	videoService service.VideoService = service.New()
// )

func main() {
	server := gin.Default()
	server.GET("/posts", func(ctx *gin.Context) {

	})
	server.Run(":8080")
}
