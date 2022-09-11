package user

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/user/request"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/repository"
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/context"

	"gorm.io/gorm"
)

type UserUseCase interface {
	AddUser(
		ctx *gorm.DB,
		params request.AddUsersReq,
	) (*model.User, error)

	ShowUser(
		ctx *gorm.DB,
		user *model.User,
	) (*model.User, error)

	DeleteUser(
		ctx *context.ContextUser,
		userId model.UserId,
	) error
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
