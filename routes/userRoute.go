package routes

import (
	"mygram/controllers"
	"mygram/database"
	"mygram/middlewares"
	"mygram/repositories"
	"mygram/services"

	"github.com/gin-gonic/gin"
)

func UserRoute(rg *gin.RouterGroup) {
	UserController := userInjection()

	userRoutes := rg.Group("/users")
	{
		userRoutes.POST("/register", UserController.Register)
		userRoutes.POST("/login", UserController.Login)
		userRoutes.PUT("/:userId", middlewares.Authentication(), middlewares.UserAuthorization(), UserController.UpdateUser)
		userRoutes.DELETE("/", middlewares.Authentication(), UserController.DeleteUser)
	}
}

func userInjection() *controllers.UserController {
	var db = database.GetDB()

	userRepo := repositories.NewUserRepo(db)
	userSvc := services.NewUserService(&userRepo)
	return controllers.NewUserController(userSvc)
}
