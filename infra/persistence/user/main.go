package user

import "github.com/htoyoda18/sample-tweet-api/domain/repository"

type UserPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &UserPersistence{}
}
