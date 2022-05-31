package main

import (
	"log"
	"mygram/controllers"
	"mygram/database"
	"mygram/repositories"
	"mygram/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const PORT = ":3636"

func main() {
	db := database.StartDB()
	r := gin.Default()

	UserController := UserInjection(db)
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", UserController.Register)
		userRoutes.POST("/login", UserController.Login)
		userRoutes.PUT("/", UserController.UpdateUser)
		userRoutes.DELETE("/", UserController.DeleteUser)
	}
	log.Println("Server is listening at port", PORT)
	r.Run(PORT)
}

func UserInjection(db *gorm.DB) *controllers.UserController {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(&userRepo)
	return controllers.NewUserController(userService)
}
