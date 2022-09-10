package injector

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/repository"
	refreshToken "github.com/htoyoda18/sample-tweet-api/golang/src/infra/persistence/refresh_token"
	"github.com/htoyoda18/sample-tweet-api/golang/src/infra/persistence/user"
)

type Infra struct {
	UserPersistence         repository.UserRepository
	RefreshTokenPersistence repository.RefreshTokenRepository
}

func NewInfra() *Infra {
	UserPersistence := user.NewUserPersistence()
	RefreshTokenPersistence := refreshToken.NewRefreshTokenPersistence()

	return &Infra{
		UserPersistence,
		RefreshTokenPersistence,
	}
}
