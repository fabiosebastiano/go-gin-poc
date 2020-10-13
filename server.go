package main

import (
	"io"
	"net/http"
	"os"

	"github.com/fabiosebastiano/go-gin-poc/controller"
	"github.com/fabiosebastiano/go-gin-poc/middlewares"
	"github.com/fabiosebastiano/go-gin-poc/repository"
	"github.com/fabiosebastiano/go-gin-poc/service"
	"github.com/gin-gonic/gin"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	videoController controller.VideoController = controller.New(videoService)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	defer videoRepository.CloseDB()

	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("./templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger())

	// Login endpoint: Autenticazione + Token Creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}

	})

	//CREIAMO GRUPPO DI ROUTES PER LE API a cui applichiamo middleware di JWT Authorization
	apiRoutes := server.Group("api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video inserted"})
			}
		})

		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video update!"})
			}
		})

		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video update!"})
			}
		})

	}

	//ALTRO GRUPPO DI ROUTES CON LE VISTE/HTML TEMPLATES
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
		//	viewRoutes.POST("/videos", nil)
	}

	// SETTA VAR ENV DALLA CONSOLE DI EB
	port := os.Getenv("PORT")
	// eb forwarda le richieste alla porta 5000 >>> USA NGNX
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
