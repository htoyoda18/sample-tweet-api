package injector

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/usecase/tweet"
	"github.com/htoyoda18/sample-tweet-api/golang/src/usecase/user"
)

type UseCase struct {
	UserUseCase  user.UserUseCase
	TweetUseCase tweet.TweetUseCase
}

func NewUseCase() *UseCase {
	infra := NewInfra()

	UserUseCase := user.NewUserUseCase(
		infra.UserPersistence,
	)
	TweetUseCase := tweet.NewTweetUseCase(
		infra.TweetPersistence,
	)

	return &UseCase{
		UserUseCase,
		TweetUseCase,
	}
}
