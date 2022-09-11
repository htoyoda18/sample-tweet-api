package tweet

import "github.com/htoyoda18/sample-tweet-api/golang/src/domain/repository"

type TweetPersistence struct{}

func NewTweetPersistence() repository.TweetRepository {
	return &TweetPersistence{}
}
