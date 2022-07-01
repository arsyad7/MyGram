package routes

import (
	"mygram/controllers"
	"mygram/database"
	"mygram/middlewares"
	"mygram/repositories"
	"mygram/services"

	"github.com/gin-gonic/gin"
)

func SocialMediaRoute(rg *gin.RouterGroup) {
	controller := socialMediaInjection()
	socmedRoutes := rg.Group("/socialmedias")
	{
		socmedRoutes.Use(middlewares.Authentication())
		socmedRoutes.POST("/", controller.CreateSocialMedia)
		socmedRoutes.GET("/", controller.GetSocialMedias)
		socmedRoutes.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controller.UpdateSocialMedia)
		socmedRoutes.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controller.DeleteSocialMedia)
	}
}

func socialMediaInjection() *controllers.SocialMediaController {
	var db = database.GetDB()
	socmedRepo := repositories.NewSocialMediaRepo(db)
	socmedSvc := services.NewSocialMediaService(&socmedRepo)
	return controllers.NewSocialMediaController(socmedSvc)
}
