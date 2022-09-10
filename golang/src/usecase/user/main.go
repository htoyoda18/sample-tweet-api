package user

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/user/request"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/repository"

	"gorm.io/gorm"
)

type UserUseCase interface {
	AddUser(
		ctx *gorm.DB,
		params request.AddUsersReq,
	) (*model.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(
	userRepository repository.UserRepository,
) UserUseCase {
	return &userUseCase{
		userRepository,
	}
}
