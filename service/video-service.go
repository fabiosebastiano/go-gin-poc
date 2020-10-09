package service

import (
	"fmt"

	"github.com/fabiosebastiano/go-gin-poc/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	fmt.Println("PRIMA SAVE ", len(service.videos))
	service.videos = append(service.videos, video)
	fmt.Println("DOPO SAVE ", len(service.videos))
	return video
}
func (service *videoService) FindAll() []entity.Video {
	fmt.Println("GET ", len(service.videos))
	return service.videos
}
