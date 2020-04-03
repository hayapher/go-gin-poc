package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-gin-poc/controller"
	"github.com/golang-gin-poc/middlewares"
	"github.com/golang-gin-poc/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	setupLogOutput()
	router := gin.New()

	router.Use(gin.Recovery(),
		middlewares.Logger(),
		middlewares.BasicAuth(),
		gindump.Dump(),
	)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK!",
		})
	})

	router.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})
	router.POST("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.Save(c))
	})

	router.Run(":8080")
}

func setupLogOutput() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}
