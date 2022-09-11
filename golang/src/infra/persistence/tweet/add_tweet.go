package tweet

import (
	"log"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"gorm.io/gorm"
)

func (TweetPersistence) AddTweet(
	db *gorm.DB,
	tweet *model.Tweet,
) (*model.Tweet, error) {
	log.Printf("AddTweet")

	if err := db.
		Create(tweet).Error; err != nil {
		log.Printf("AddTweet Erorr %s", err)
		return nil, err
	}

	return tweet, nil
}
