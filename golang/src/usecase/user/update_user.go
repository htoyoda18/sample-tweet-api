package user

import (
	"errors"
	"log"

	"github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/user/request"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/context"
)

func (uu userUseCase) UpdateUser(
	ctx *context.ContextUser,
	userId model.UserId,
	params request.UpdateUserReq,
) (*model.User, error) {
	db := ctx.DB

	oldUser, err := uu.userRepository.ShowUser(db, &model.User{
		ID: userId,
	})
	if err != nil {
		log.Println("Erorr UpdateUser", err)
		return nil, err
	}

	if oldUser.ID != ctx.UserId {
		err := errors.New("fail to other user")
		log.Println("Erorr UpdateUser", err)
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
		log.Println("Erorr UpdateUser", UserErr)
		return nil, UserErr
	}

	return newUser, nil
}
