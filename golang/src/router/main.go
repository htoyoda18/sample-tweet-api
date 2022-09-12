package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/htoyoda18/sample-tweet-api/golang/src/docs" // swag initで生成したdocs。
	"github.com/htoyoda18/sample-tweet-api/golang/src/injector"
	"github.com/htoyoda18/sample-tweet-api/golang/src/router/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	handler := injector.NewHandler()

	r := gin.Default()

	r.Use(middleware.SetHeaderHandler())
	r.Use(middleware.LoggerMiddleware())

	v1 := r.Group("v1")
	{
		v1.POST("/signup", handler.UserHandler.AddUser)
		v1.POST("/login", handler.LoginHandler.Login)

		user := v1.Group("user")
		user.Use(middleware.AuthMiddleware())
		{
			user.GET("/:id", handler.UserHandler.ShowUser)
			user.DELETE("/:id", handler.UserHandler.DeleteUser)
			user.PATCH("/:id", handler.UserHandler.UpdateUser)
		}

		tweet := v1.Group("tweet")
		tweet.Use(middleware.AuthMiddleware())
		{
			tweet.POST("", handler.TweetHandler.AddTweet)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
