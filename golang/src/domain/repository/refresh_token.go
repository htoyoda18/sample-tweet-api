package repository

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	AddRefreshToken(
		db *gorm.DB,
		refreshToken *model.RefreshToken,
	) (*model.RefreshToken, error)
}
