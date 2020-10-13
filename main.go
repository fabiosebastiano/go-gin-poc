package main

import (
	"io"
	"os"

	"github.com/fabiosebastiano/go-gin-poc/api"
	"github.com/fabiosebastiano/go-gin-poc/controller"
	"github.com/fabiosebastiano/go-gin-poc/middlewares"
	"github.com/fabiosebastiano/go-gin-poc/repository"
	"github.com/fabiosebastiano/go-gin-poc/service"

	"github.com/fabiosebastiano/go-gin-poc/docs" // swagger generated files
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {

	// Swagger metadata
	docs.SwaggerInfo.Title = "GO GIN TUTORIAL API "
	docs.SwaggerInfo.Description = " Fabio Sebastiano - Video API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	defer videoRepository.CloseDB()

	server := gin.Default()

	videoAPI := api.NewVideoAPI(loginController, videoController)

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", videoAPI.Authenticate)
		}
		videos := apiRoutes.Group("/videos", middlewares.AuthorizeJWT())
		{
			videos.GET("", videoAPI.GetVideos)
			videos.POST("", videoAPI.CreateVideo)
			videos.PUT(":id", videoAPI.UpdateVideo)
			videos.DELETE(":id", videoAPI.DeleteVideo)
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// SETTA VAR ENV DALLA CONSOLE DI EB
	port := os.Getenv("PORT")
	// eb forwarda le richieste alla porta 5000 >>> USA NGNX
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
