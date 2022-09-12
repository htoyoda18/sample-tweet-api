package user

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/shared/logger"
	"gorm.io/gorm"
)

func (UserPersistence) AddUser(
	db *gorm.DB,
	user *model.User,
) (*model.User, error) {
	logger.Info("AddUser")

	if err := db.
		Create(user).Error; err != nil {
		logger.Error("AddUser", err)
		return nil, err
	}

	return user, nil
}
