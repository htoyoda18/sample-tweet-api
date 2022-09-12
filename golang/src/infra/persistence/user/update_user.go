package user

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/shared/logger"
	"gorm.io/gorm"
)

func (UserPersistence) UpdateUser(
	db *gorm.DB,
	userId model.UserId,
	where *model.User,
) (*model.User, error) {
	logger.Info("UpdateUser")

	user := &model.User{ID: userId}

	if err := db.
		Model(user).
		Updates(where).Error; err != nil {
		logger.Error("UpdateUser", err)
		return nil, err
	}

	return user, nil
}
