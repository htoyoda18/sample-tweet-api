package repository

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(
		db *gorm.DB,
		user *model.User,
	) (*model.User, error)

	ShowUser(
		db *gorm.DB,
		user *model.User,
	) (*model.User, error)

	DeleteUser(
		db *gorm.DB,
		user *model.User,
	) error
}
