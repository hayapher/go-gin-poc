package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-gin-poc/entity"
	"github.com/golang-gin-poc/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(c *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (ctl *controller) FindAll() []entity.Video {
	return ctl.service.FindAll()
}
func (ctl *controller) Save(c *gin.Context) entity.Video {
	var video entity.Video
	c.BindJSON(&video)
	ctl.service.Save(video)
	return video
}
