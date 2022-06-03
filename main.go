package main

import (
	"log"
	"mygram/controllers"
	"mygram/database"
	"mygram/middlewares"
	"mygram/repositories"
	"mygram/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const PORT = ":3636"

var DB = database.StartDB()

func main() {
	r := gin.Default()

	UserController := UserInjection(DB)
	PhotoController := PhotoInjection(DB)
	CommentController := CommentInjection(DB)
	SocialMediaController := SocialMediaInjection(DB)

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", UserController.Register)
		userRoutes.POST("/login", UserController.Login)
		userRoutes.PUT("/:userId", middlewares.Authentication(), middlewares.UserAuthorization(), UserController.UpdateUser)
		userRoutes.DELETE("/", middlewares.Authentication(), UserController.DeleteUser)
	}
	photoRoutes := r.Group("/photos")
	{
		photoRoutes.Use(middlewares.Authentication())
		photoRoutes.POST("/", PhotoController.CreatePhoto)
		photoRoutes.GET("/", PhotoController.GetPhotos)
		photoRoutes.PUT("/:photoId", middlewares.PhotoAuthorization(), PhotoController.UpdatePhoto)
		photoRoutes.DELETE("/:photoId", middlewares.PhotoAuthorization(), PhotoController.DeletePhoto)
	}
	commentRoutes := r.Group("/comments")
	{
		commentRoutes.Use(middlewares.Authentication())
		commentRoutes.POST("/", CommentController.CreateComment)
		commentRoutes.GET("/", CommentController.GetComents)
		commentRoutes.PUT("/:commentId", middlewares.CommentAuthorization(), CommentController.UpdateComment)
		commentRoutes.DELETE("/:commentId", middlewares.CommentAuthorization(), CommentController.DeleteComment)
	}
	socmedRoutes := r.Group("/socialmedias")
	{
		socmedRoutes.Use(middlewares.Authentication())
		socmedRoutes.POST("/", SocialMediaController.CreateSocialMedia)
		socmedRoutes.GET("/", SocialMediaController.GetSocialMedias)
		socmedRoutes.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), SocialMediaController.UpdateSocialMedia)
		socmedRoutes.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), SocialMediaController.DeleteSocialMedia)
	}

	log.Println("Server is listening at port", PORT)
	r.Run(PORT)
}

func UserInjection(db *gorm.DB) *controllers.UserController {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(&userRepo)
	return controllers.NewUserController(userService)
}

func PhotoInjection(db *gorm.DB) *controllers.PhotoController {
	photoRepo := repositories.NewPhotoRepo(db)
	photoService := services.NewPhotoService(&photoRepo)
	return controllers.NewPhotoController(photoService)
}

func CommentInjection(db *gorm.DB) *controllers.CommentController {
	commentRepo := repositories.NewCommentRepo(db)
	commentService := services.NewCommentService(&commentRepo)
	return controllers.NewCommentController(commentService)
}

func SocialMediaInjection(db *gorm.DB) *controllers.SocialMediaController {
	socialMediaRepo := repositories.NewSocialMediaRepo(db)
	socialMediaService := services.NewSocialMediaService(&socialMediaRepo)
	return controllers.NewSocialMediaController(socialMediaService)
}
