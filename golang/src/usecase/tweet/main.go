package tweet

import (
	request "github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/tweet/request"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/repository"
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/context"
)

type TweetUseCase interface {
	AddTweet(
		ctx *context.ContextUser,
		params request.AddTweetReq,
	) (*model.Tweet, error)
}

type tweetUseCase struct {
	tweetRepository repository.TweetRepository
}

func NewTweetUseCase(
	tweetRepository repository.TweetRepository,
) TweetUseCase {
	return &tweetUseCase{
		tweetRepository,
	}
}
