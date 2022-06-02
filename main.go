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
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", UserController.Register)
		userRoutes.POST("/login", UserController.Login)
		userRoutes.PUT("/", middlewares.Authentication(), UserController.UpdateUser)
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
