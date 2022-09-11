package user

import (
	"log"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"gorm.io/gorm"
)

func (UserPersistence) DeleteUser(
	db *gorm.DB,
	where *model.User,
) error {
	log.Printf("DeleteUser")

	var user model.User

	db.Where(where).Delete(&user)

	return nil
}
