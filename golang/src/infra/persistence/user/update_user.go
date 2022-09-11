package user

import (
	"log"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"gorm.io/gorm"
)

func (UserPersistence) UpdateUser(
	db *gorm.DB,
	userId model.UserId,
	where *model.User,
) (*model.User, error) {
	log.Printf("UpdateUser")

	user := &model.User{ID: userId}

	if err := db.
		Model(user).
		Updates(where).Error; err != nil {
		log.Printf("UpdateUser Erorr %s", err)
		return nil, err
	}

	return user, nil
}
