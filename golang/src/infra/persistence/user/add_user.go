package user

import (
	"log"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"gorm.io/gorm"
)

func (UserPersistence) AddUser(
	db *gorm.DB,
	user *model.User,
) (*model.User, error) {
	log.Printf("AddUser")

	if err := db.
		Create(user).Error; err != nil {
		log.Printf("AddUser Erorr %s", err)
		return nil, err
	}

	return user, nil
}
