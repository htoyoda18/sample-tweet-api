package injector

import (
	"github.com/htoyoda18/sample-tweet-api/controller/handler/user"
)

type Handler struct {
	UserHandler user.UserHandler
}

func NewHandler() *Handler {
	usecase := NewUseCase()
	UserHandler := user.NewUserHandler(usecase.UserUseCase)

	return &Handler{
		UserHandler,
	}
}
