package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/htoyoda18/sample-tweet-api/controller/handler/user/request"
	"github.com/htoyoda18/sample-tweet-api/service/context"
	"github.com/htoyoda18/sample-tweet-api/usecase/user"
)

type UserHandler interface {
	AddUser(*gin.Context)
}

type userHandler struct {
	ctx         context.ContextApi
	userUseCase user.UserUseCase
}

func NewUserHandler(
	ctx context.ContextApi,
	userUseCase user.UserUseCase,
) UserHandler {
	return &userHandler{
		ctx,
		userUseCase,
	}
}

func (uh userHandler) AddUser(c *gin.Context) {
	log.Printf("AddUser")

	var params request.AddUsersReq

	if err := c.ShouldBind(&params); err != nil {
		log.Printf("Erorr AddUser")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx, ctxErr := uh.ctx.Context(c)
	if ctxErr != nil {
		log.Printf("Erorr AddUser")
		c.AbortWithError(http.StatusBadRequest, ctxErr)
		return
	}

	user, err := uh.userUseCase.AddUser(ctx, params)
	if err != nil {
		log.Printf("Erorr AddUser")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(200, user)
}
