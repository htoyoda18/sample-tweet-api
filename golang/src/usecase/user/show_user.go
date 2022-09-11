package user

import (
	"log"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"gorm.io/gorm"
)

func (uu userUseCase) ShowUser(ctx *gorm.DB, user *model.User) (*model.User, error) {
	log.Println("ShowUser")

	getUser, err := uu.userRepository.ShowUser(ctx, user)

	if err != nil {
		log.Println("ShowUser Error: ", err)
		return nil, err
	}

	return getUser, nil
}
