package tweet

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	request "github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/tweet/request"
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/context"
	"github.com/htoyoda18/sample-tweet-api/golang/src/usecase/tweet"
)

type TweetHandler interface {
	AddTweet(*gin.Context)
}

type tweetHandler struct {
	ctx          context.ContextApi
	tweetUseCase tweet.TweetUseCase
}

func NewTweetHandler(
	ctx context.ContextApi,
	tweetUseCase tweet.TweetUseCase,
) TweetHandler {
	return &tweetHandler{
		ctx,
		tweetUseCase,
	}
}

// AddTweet ツイート投稿のAPI
// @Summary ツイート投稿のAPI
// @description ツイート投稿のAPI
// @param request body request.AddTweetReq true "request body"
// @Success 200
// @Failure 400
// @Failure 500
// @router /tweet [POST]
func (th tweetHandler) AddTweet(c *gin.Context) {
	log.Printf("AddTweet")

	var params request.AddTweetReq

	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("Erorr AddTweet")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx, ctxErr := th.ctx.ContextUser(c)

	if ctxErr != nil {
		log.Printf("Erorr AddTweet")
		c.AbortWithError(http.StatusBadRequest, ctxErr)
		return
	}

	tweet, err := th.tweetUseCase.AddTweet(ctx, params)
	if err != nil {
		log.Printf("Erorr AddTweet")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(200, tweet)
}
