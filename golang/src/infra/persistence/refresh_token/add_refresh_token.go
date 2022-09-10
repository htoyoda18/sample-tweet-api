package refreshToken

import (
	"log"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (RefreshTokenPersistence) AddRefreshToken(
	db *gorm.DB,
	refreshToken *model.RefreshToken,
) (*model.RefreshToken, error) {
	log.Println("AddRefreshToken")

	if err := db.
		Omit(clause.Associations).
		Create(refreshToken).Error; err != nil {
		log.Println("AddRefreshToken Error: ", err)
		return nil, err
	}

	return refreshToken, nil
}
