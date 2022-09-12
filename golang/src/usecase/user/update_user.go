package user

import (
	"errors"

	"github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/user/request"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/context"
	"github.com/htoyoda18/sample-tweet-api/golang/src/shared/logger"
)

func (uu userUseCase) UpdateUser(
	ctx *context.ContextUser,
	userId model.UserId,
	params request.UpdateUserReq,
) (*model.User, error) {
	logger.Info("UpdateUser")
	db := ctx.DB

	oldUser, err := uu.userRepository.ShowUser(db, &model.User{
		ID: userId,
	})
	if err != nil {
		logger.Error("UpdateUser", err)
		return nil, err
	}

	if oldUser.ID != ctx.UserId {
		err := errors.New("fail to other user")
		logger.Error("UpdateUser", err)
		return nil, err
	}

	newUser, UserErr := uu.userRepository.UpdateUser(db,
		userId,
		&model.User{
			Email:    params.Email,
			Name:     params.Name,
			Password: params.Password,
			Icon:     params.Icon,
		})
	if UserErr != nil {
		logger.Error("UpdateUser", UserErr)
		return nil, UserErr
	}

	return newUser, nil
}
