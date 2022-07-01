package routes

import (
	"mygram/controllers"
	"mygram/database"
	"mygram/middlewares"
	"mygram/repositories"
	"mygram/services"

	"github.com/gin-gonic/gin"
)

func PhotoRoute(rg *gin.RouterGroup) {
	controller := photoInjection()

	photoRoutes := rg.Group("/photos")
	{
		photoRoutes.Use(middlewares.Authentication())
		photoRoutes.POST("/", controller.CreatePhoto)
		photoRoutes.GET("/", controller.GetPhotos)
		photoRoutes.PUT("/:photoId", middlewares.PhotoAuthorization(), controller.UpdatePhoto)
		photoRoutes.DELETE("/:photoId", middlewares.PhotoAuthorization(), controller.DeletePhoto)
	}
}

func photoInjection() *controllers.PhotoController {
	var db = database.GetDB()
	photoRepo := repositories.NewPhotoRepo(db)
	photoSvc := services.NewPhotoService(&photoRepo)
	return controllers.NewPhotoController(photoSvc)
}
