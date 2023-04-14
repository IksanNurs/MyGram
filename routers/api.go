package router

import (
	"finalproject_mygram/controllers"
	"finalproject_mygram/database"
	"finalproject_mygram/middlewares"
	"finalproject_mygram/repository"
	"finalproject_mygram/service"
	"net/http"

	_ "finalproject_mygram/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	COMMENT     = 0
	PHOTO       = 1
	SOCIALMEDIA = 2
)

// @title MyGram API
// version 1.0
// @description This is a sample service for managing books
// termOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkader@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html
// @host localhost:8086
// @BasePath/
func StartApp() *gin.Engine {
	db := database.GetDB()
	//=====
	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	commentController := controllers.NewCommentController(commentService)
	//===
	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	photoController := controllers.NewPhotoController(photoService)
	//===
	socialmediaRepository := repository.NewSocialMediaRepository(db)
	socialmediaService := service.NewSocialMediaService(socialmediaRepository)
	socialmediaController := controllers.NewSocialMediaController(socialmediaService)

	r := gin.Default()
	r.GET("/coba", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "berhasil")
	})
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", commentController.CreateComment)
		commentRouter.GET("/", commentController.GetAllComment)
		commentRouter.GET("/:commentId", commentController.GetOneComment)
		commentRouter.PUT("/:commentId", middlewares.Authorization(&COMMENT), commentController.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.Authorization(&COMMENT), commentController.DeleteComment)
	}
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", photoController.CreatePhoto)
		photoRouter.GET("/", photoController.GetAllPhoto)
		photoRouter.GET("/:photoId", photoController.GetOnePhoto)
		photoRouter.PUT("/:photoId", middlewares.Authorization(&PHOTO), photoController.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.Authorization(&PHOTO), photoController.DeletePhoto)
	}
	socialmediaRouter := r.Group("/socialmedias")
	{
		socialmediaRouter.Use(middlewares.Authentication())
		socialmediaRouter.POST("/", socialmediaController.CreateSocialMedia)
		socialmediaRouter.GET("/", socialmediaController.GetAllSocialMedia)
		socialmediaRouter.GET("/:socialmediaId", socialmediaController.GetOneSocialMedia)
		socialmediaRouter.PUT("/:socialmediaId", middlewares.Authorization(&SOCIALMEDIA), socialmediaController.UpdateSocialMedia)
		socialmediaRouter.DELETE("/:socialmediaId", middlewares.Authorization(&SOCIALMEDIA), socialmediaController.DeleteSocialMedia)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
