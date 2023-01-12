package main

import (
	"example/golang_gin/controller"
	"example/golang_gin/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()
	videoContorller controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message" : "OK!!",
	// 	})
	// })

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoContorller.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoContorller.Save(ctx))
	})

	server.Run(":8080")
}