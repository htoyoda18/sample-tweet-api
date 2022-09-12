package user

import (
	"errors"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/context"
	"github.com/htoyoda18/sample-tweet-api/golang/src/shared/logger"
)

func (uu userUseCase) DeleteUser(
	ctx *context.ContextUser,
	userId model.UserId,
) error {
	logger.Info("DeleteUser userId by ", userId)

	db := ctx.DB

	user, err := uu.userRepository.ShowUser(db, &model.User{
		ID: userId,
	})
	if err != nil {
		logger.Error("DeleteUser", err)
		return err
	}

	if user.ID != ctx.UserId {
		err := errors.New("fail to other user")
		logger.Error("DeleteUser", err)
		return err
	}

	UserErr := uu.userRepository.DeleteUser(db, &model.User{
		ID: userId,
	})
	if UserErr != nil {
		logger.Error("DeleteUser", UserErr)
		return err
	}

	return nil
}
