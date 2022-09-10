package refreshToken

import "github.com/htoyoda18/sample-tweet-api/golang/src/domain/repository"

type RefreshTokenPersistence struct{}

func NewRefreshTokenPersistence() repository.RefreshTokenRepository {
	return &RefreshTokenPersistence{}
}
