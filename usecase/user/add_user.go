package user

import (
	"errors"
	"log"

	"github.com/htoyoda18/sample-tweet-api/controller/handler/user/request"
	"github.com/htoyoda18/sample-tweet-api/domain/model"
	"gorm.io/gorm"
)

func (uu userUseCase) AddUser(ctx *gorm.DB, params request.AddUsersReq) (*model.User, error) {
	log.Printf("AddUser")

	// メールアドレスでユーザーを取得
	selectUser, _ := uu.userRepository.ShowUser(&model.User{
		Email: params.Email,
	})

	// 既に登録しているアドレスの場合
	if selectUser != nil {
		err := errors.New("email is already exists")
		log.Printf("AddUser email is already exists")
		return nil, err
	}

	// ユーザーを追加
	user, errCreate := uu.userRepository.AddUser(
		&model.User{
			Name:     params.Name,
			Email:    params.Email,
			Password: params.Password,
		},
	)

	// ユーザーの追加が問題無いかチェック
	if errCreate != nil {
		log.Printf("AddUser %s", errCreate)
		return nil, errCreate
	}

	return user, nil
}
