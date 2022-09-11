package injector

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/login"
	"github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/user"
)

type Handler struct {
	UserHandler  user.UserHandler
	LoginHandler login.LoginHandler
}

func NewHandler() *Handler {
	usecase := NewUseCase()
	service := NewService()

	UserHandler := user.NewUserHandler(
		service.ApiContext,
		usecase.UserUseCase,
	)
	LoginHandler := login.NewLoginHandler(
		service.ApiContext,
		service.AuthCore,
		usecase.UserUseCase,
	)

	return &Handler{
		UserHandler,
		LoginHandler,
	}
}
