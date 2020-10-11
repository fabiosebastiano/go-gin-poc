package main

import (
	"io"
	"net/http"
	"os"

	"github.com/fabiosebastiano/go-gin-poc/controller"
	"github.com/fabiosebastiano/go-gin-poc/middlewares"
	"github.com/fabiosebastiano/go-gin-poc/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger()) //	middlewares.BasicAuth(),
	//	, gindump.Dump()

	//CREIAMO GRUPPO DI ROUTES PER LE API
	apiRoutes := server.Group("api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid"})
			}
		})
	}

	//ALTRO GRUPPO DI ROUTES CON LE VISTE/HTML TEMPLATES
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
		//	viewRoutes.POST("/videos", nil)
	}

	server.Run(":8080")
}
