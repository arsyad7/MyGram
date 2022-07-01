package routes

import (
	"mygram/controllers"
	"mygram/database"
	"mygram/middlewares"
	"mygram/repositories"
	"mygram/services"

	"github.com/gin-gonic/gin"
)

func CommentRoute(rg *gin.RouterGroup) {
	controller := commentInjection()
	commentRoutes := rg.Group("/comments")
	{
		commentRoutes.Use(middlewares.Authentication())
		commentRoutes.POST("/", controller.CreateComment)
		commentRoutes.GET("/", controller.GetComents)
		commentRoutes.PUT("/:commentId", middlewares.CommentAuthorization(), controller.UpdateComment)
		commentRoutes.DELETE("/:commentId", middlewares.CommentAuthorization(), controller.DeleteComment)
	}
}

func commentInjection() *controllers.CommentController {
	var db = database.GetDB()
	commentRepo := repositories.NewCommentRepo(db)
	commentSvc := services.NewCommentService(&commentRepo)
	return controllers.NewCommentController(commentSvc)
}
