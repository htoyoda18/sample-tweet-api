package tweet

import (
	"log"

	request "github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/tweet/request"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/context"
)

func (tu tweetUseCase) AddTweet(
	ctx *context.ContextUser,
	params request.AddTweetReq,
) (*model.Tweet, error) {
	log.Printf("AddTweet")

	tweet, createErr := tu.tweetRepository.AddTweet(ctx.DB,
		&model.Tweet{
			Text:   params.Text,
			UserId: ctx.UserId,
		},
	)

	if createErr != nil {
		log.Printf("AddTweet %s", createErr)
		return nil, createErr
	}

	return tweet, nil
}
