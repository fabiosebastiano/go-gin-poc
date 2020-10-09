package controller

import (
	"fmt"

	"github.com/fabiosebastiano/go-gin-poc/entity"
	"github.com/fabiosebastiano/go-gin-poc/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	fmt.Println(video)
	c.service.Save(video)
	return video
}
