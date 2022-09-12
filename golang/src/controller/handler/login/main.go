package login

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/login/request"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/context"
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/core"
	"github.com/htoyoda18/sample-tweet-api/golang/src/usecase/user"
)

type LoginHandler interface {
	Login(c *gin.Context)
}

type loginHandler struct {
	ctx         context.ContextApi
	authCore    core.AuthCore
	userUseCase user.UserUseCase
}

func NewLoginHandler(
	ctx context.ContextApi,
	authCore core.AuthCore,
	userUseCase user.UserUseCase,
) LoginHandler {
	return &loginHandler{
		ctx,
		authCore,
		userUseCase,
	}
}

// Login ログイン用のAPI
// @Summary ログイン用のAPI
// @description ログイン用のAPI
// @param request body request.LoginReq true "request body"
// @Success 200
// @Failure 400
// @Failure 500
// @router /login [POST]
func (lh loginHandler) Login(c *gin.Context) {
	log.Printf("Login")

	var params request.LoginReq

	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("Erorr AddUser")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx, ctxErr := lh.ctx.Context(c)
	if ctxErr != nil {
		log.Println("Erorr Login", ctxErr)
		c.AbortWithError(http.StatusBadRequest, ctxErr)
		return
	}

	user, err := lh.userUseCase.ShowUser(ctx, &model.User{
		Email:    params.Email,
		Password: params.Password,
	})

	if err != nil {
		log.Println("Erorr Login", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	jwt := lh.authCore.NewJwt(&model.User{
		ID: user.ID,
	})

	c.SetCookie("jwt", jwt, 3600, "/", "localhost", false, true)

	c.Status(200)
}
