package user

import (
	"log"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	sharedDB "github.com/htoyoda18/sample-tweet-api/golang/src/shared/db"
	"gorm.io/gorm"
)

func (UserPersistence) ShowUser(
	db *gorm.DB,
	where *model.User,
) (*model.User, error) {
	log.Printf("ShowUser")

	user := &model.User{}

	gormDB, sqlDB, _ := sharedDB.InitDB()
	defer sqlDB.Close()

	if err := gormDB.
		Where(where).
		First(user).Error; err != nil {
		log.Printf("ShowUser Erorr %s", err)
	}

	return user, nil
}
