package repository

import (
	"github.com/fabiosebastiano/go-gin-poc/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Fallito mentre aprivo la connessione")
	}

	db.AutoMigrate(&entity.Video{}, &entity.Person{})

	return &database{
		connection: db,
	}
}

func (database *database) CloseDB() {
	err := database.connection.Close()
	if err != nil {
		panic("Fallito durante chiusura della connessione")
	}
}

func (database *database) Save(video entity.Video) {
	database.connection.Create(&video)
}
func (database *database) Update(video entity.Video) {
	database.connection.Save(&video)
}
func (database *database) Delete(video entity.Video) {
	database.connection.Delete(&video)
}
func (database *database) FindAll() []entity.Video {
	var videos []entity.Video
	database.connection.Set("gorm:auto_preload", true).Find(&videos)

	return videos
}
