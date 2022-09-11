package router

import (
	"github.com/gin-gonic/gin"
	"github.com/htoyoda18/sample-tweet-api/golang/src/injector"
	"github.com/htoyoda18/sample-tweet-api/golang/src/router/middleware"
)

func SetupRouter() *gin.Engine {
	handler := injector.NewHandler()

	r := gin.Default()

	r.Use(middleware.SetHeaderHandler())

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
	}

	return r
}
