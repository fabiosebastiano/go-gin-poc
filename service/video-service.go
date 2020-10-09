package service

import (
	"github.com/fabiosebastiano/go-gin-poc/entity"
)

type VideoService interface {
	Save(entity.Video) error
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) error {
	service.videos = append(service.videos, video)
	return nil
}
func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
