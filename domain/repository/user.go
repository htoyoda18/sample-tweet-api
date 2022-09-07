package repository

import "github.com/htoyoda18/sample-tweet-api/domain/model"

type UserRepository interface {
	AddUser(
		user *model.User,
	) (*model.User, error)

	ShowUser(
		user *model.User,
	) (*model.User, error)
}
