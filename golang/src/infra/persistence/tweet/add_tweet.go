package tweet

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/shared/logger"
	"gorm.io/gorm"
)

func (TweetPersistence) AddTweet(
	db *gorm.DB,
	tweet *model.Tweet,
) (*model.Tweet, error) {
	logger.Info("AddTweet")

	if err := db.
		Create(tweet).Error; err != nil {
		logger.Info("AddTweet", err)
		return nil, err
	}

	return tweet, nil
}
