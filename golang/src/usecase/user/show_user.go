package user

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/shared/logger"
	"gorm.io/gorm"
)

func (uu userUseCase) ShowUser(ctx *gorm.DB, user *model.User) (*model.User, error) {
	logger.Info("ShowUser")

	getUser, err := uu.userRepository.ShowUser(ctx, user)

	if err != nil {
		logger.Error("ShowUser", err)
		return nil, err
	}

	return getUser, nil
}
