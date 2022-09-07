package user

import (
	"github.com/gin-gonic/gin"
	"github.com/htoyoda18/sample-tweet-api/usecase/user"
)

type UserHandler interface {
	AddUser(*gin.Context)
}

type userHandler struct {
	userUseCase user.UserUseCase
}

func NewUserHandler(
	userUseCase user.UserUseCase,
) UserHandler {
	return &userHandler{
		userUseCase,
	}
}

func (uh userHandler) AddUser(c *gin.Context) {

}
