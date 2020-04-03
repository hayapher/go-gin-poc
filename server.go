package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-gin-poc/controller"
	"github.com/golang-gin-poc/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	router := gin.New()

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
