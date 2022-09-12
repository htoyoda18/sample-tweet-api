package user

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/shared/logger"
	"gorm.io/gorm"
)

func (UserPersistence) DeleteUser(
	db *gorm.DB,
	where *model.User,
) error {
	logger.Info("DeleteUser")

	var user model.User

	db.Where(where).Delete(&user)

	return nil
}
