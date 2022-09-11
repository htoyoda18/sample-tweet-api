package injector

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/login"
	"github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/tweet"
	"github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/user"
)

type Handler struct {
	UserHandler  user.UserHandler
	LoginHandler login.LoginHandler
	TweetHandler tweet.TweetHandler
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
	TweetHandler := tweet.NewTweetHandler(
		service.ApiContext,
		usecase.TweetUseCase,
	)

	return &Handler{
		UserHandler,
		LoginHandler,
		TweetHandler,
	}
}
