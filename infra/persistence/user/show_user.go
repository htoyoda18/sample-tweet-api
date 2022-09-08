package user

import (
	"log"

	"github.com/htoyoda18/sample-tweet-api/domain/model"
	"gorm.io/gorm"
)

func (UserPersistence) ShowUser(
	db *gorm.DB,
	where *model.User,
) (*model.User, error) {
	log.Printf("ShowUser")

	user := &model.User{}

	if err := db.
		Where(where).
		First(user).Error; err != nil {
		log.Printf("ShowUser Erorr %s", err)
		return nil, err
	}

	return user, nil
}
