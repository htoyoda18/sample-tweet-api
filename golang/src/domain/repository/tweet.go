package repository

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"gorm.io/gorm"
)

type TweetRepository interface {
	AddTweet(
		db *gorm.DB,
		tweet *model.Tweet,
	) (*model.Tweet, error)
}
