package core

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/repository"
)

type AuthCore interface {
	NewJwt(user *model.User) string
}

type authCore struct {
	userRepository         repository.UserRepository
	refreshTokenRepository repository.RefreshTokenRepository
}

func NewAuthCore(
	userRepository repository.UserRepository,
	refreshTokenRepository repository.RefreshTokenRepository,
) AuthCore {
	return &authCore{
		userRepository,
		refreshTokenRepository,
	}
}
