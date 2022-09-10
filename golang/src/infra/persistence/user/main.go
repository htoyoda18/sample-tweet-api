package user

import "github.com/htoyoda18/sample-tweet-api/golang/src/domain/repository"

type UserPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &UserPersistence{}
}
