package user

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/shared/logger"
	"gorm.io/gorm"
)

func (UserPersistence) ShowUser(
	db *gorm.DB,
	where *model.User,
) (*model.User, error) {
	logger.Info("ShowUser")

	user := &model.User{}

	if err := db.
		Where(where).
		First(user).Error; err != nil {
		logger.Error("ShowUser", err)
		return nil, err
	}

	return user, nil
}
