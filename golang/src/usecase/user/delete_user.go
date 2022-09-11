package user

import (
	"errors"
	"log"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/context"
)

func (uu userUseCase) DeleteUser(
	ctx *context.ContextUser,
	userId model.UserId,
) error {
	log.Printf("DeleteUser userId by %d", userId)

	db := ctx.DB

	user, err := uu.userRepository.ShowUser(db, &model.User{
		ID: userId,
	})
	if err != nil {
		log.Println("Erorr DeleteUser", err)
		return err
	}

	if user.ID != ctx.UserId {
		err := errors.New("fail to other user")
		log.Println("Erorr DeleteUser", err)
		return err
	}

	UserErr := uu.userRepository.DeleteUser(db, &model.User{
		ID: userId,
	})
	if UserErr != nil {
		log.Println("Erorr DeleteUser", UserErr)
		return err
	}

	return nil
}
