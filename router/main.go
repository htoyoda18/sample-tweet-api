package router

import (
	"github.com/gin-gonic/gin"
	"github.com/htoyoda18/sample-tweet-api/injector"
)

func SetupRouter() *gin.Engine {
	handler := injector.NewHandler()

	r := gin.Default()

	v1 := r.Group("v1")
	{
		v1.POST("/signup", handler.UserHandler.AddUser)
	}

	return r
}
