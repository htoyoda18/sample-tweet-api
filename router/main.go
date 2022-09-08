package router

import (
	"github.com/gin-gonic/gin"
	"github.com/htoyoda18/sample-tweet-api/injector"
	"github.com/htoyoda18/sample-tweet-api/router/middleware"
)

func SetupRouter() *gin.Engine {
	handler := injector.NewHandler()

	r := gin.Default()

	r.Use(middleware.SetHeaderHandler())

	v1 := r.Group("v1")
	{
		v1.POST("/signup", handler.UserHandler.AddUser)
	}

	return r
}
