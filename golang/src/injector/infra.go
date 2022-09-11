package injector

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/repository"
	"github.com/htoyoda18/sample-tweet-api/golang/src/infra/persistence/tweet"
	"github.com/htoyoda18/sample-tweet-api/golang/src/infra/persistence/user"
)

type Infra struct {
	UserPersistence  repository.UserRepository
	TweetPersistence repository.TweetRepository
}

func NewInfra() *Infra {
	UserPersistence := user.NewUserPersistence()
	TweetPersistence := tweet.NewTweetPersistence()

	return &Infra{
		UserPersistence,
		TweetPersistence,
	}
}
