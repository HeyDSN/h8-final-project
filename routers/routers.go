package routers

import (
	"final-project/controllers"
	"final-project/database"
	"final-project/middleware"
	"final-project/repository"
	"final-project/services"

	"github.com/gin-gonic/gin"
)

// @securityDefinitions.basic BasicAuth
func StartServer() *gin.Engine {
	db := database.GetDB()
	router := gin.Default()

	// User routes
	userRepo := repository.UserRepo{Conn: db}
	userSvc := services.UserSvc{UserRepo: userRepo}
	userCtrl := controllers.UserController{UserSvc: userSvc}
	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userCtrl.UserRegistration)
		userRouter.POST("/login", userCtrl.UserLogin)
		userRouter.PUT("/", middleware.Authenticated(), userCtrl.UpdateUser)
		userRouter.DELETE("/", middleware.Authenticated(), userCtrl.DeleteUser)
	}

	// Photo routes
	photoRepo := repository.PhotoRepo{Conn: db}
	photoSvc := services.PhotoSvc{PhotoRepo: photoRepo}
	photoCtrl := controllers.PhotoController{PhotoSvc: photoSvc}
	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middleware.Authenticated())
		photoRouter.POST("/", photoCtrl.PostPhoto)
		photoRouter.GET("/", photoCtrl.GetPhotos)
		photoRouter.PUT("/:id", photoCtrl.UpdatePhoto)
		photoRouter.DELETE("/:id", photoCtrl.DeletePhoto)
	}

	// Comment routes
	commentRepo := repository.CommentRepo{Conn: db}
	commentSvc := services.CommentSvc{CommentRepo: commentRepo}
	commentCtrl := controllers.CommentController{CommentSvc: commentSvc}
	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middleware.Authenticated())
		commentRouter.POST("/", commentCtrl.PostComment)
		commentRouter.GET("/", commentCtrl.GetComments)
		commentRouter.PUT("/:id", commentCtrl.UpdateComment)
		commentRouter.DELETE("/:id", commentCtrl.DeleteComment)
	}

	// Social Media routes
	socialMediaRepo := repository.SocialMediaRepo{Conn: db}
	socialMediaSvc := services.SocialMediaSvc{SocialMediaRepo: socialMediaRepo}
	socialMediaCtrl := controllers.SocialMediaController{SocialMediaSvc: socialMediaSvc}
	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authenticated())
		socialMediaRouter.POST("/", socialMediaCtrl.AddSocialMedia)
		socialMediaRouter.GET("/", socialMediaCtrl.GetSocialMedias)
		socialMediaRouter.PUT("/:id", socialMediaCtrl.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", socialMediaCtrl.DeleteSocialMedia)
	}

	return router
}
