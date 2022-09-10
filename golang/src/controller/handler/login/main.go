package login

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/login/request"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/context"
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/core"
)

type LoginHandler interface {
	Login(c *gin.Context)
}

type loginHandler struct {
	ctx      context.ContextApi
	authCore core.AuthCore
}

func NewLoginHandler(
	ctx context.ContextApi,
	authCore core.AuthCore,
) LoginHandler {
	return &loginHandler{
		ctx,
		authCore,
	}
}

func (lh loginHandler) Login(c *gin.Context) {
	log.Printf("Login")

	var params request.LoginReq

	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("Erorr AddUser")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, ctxErr := lh.ctx.Context(c)
	if ctxErr != nil {
		log.Println("Erorr Login", ctxErr)
		c.AbortWithError(http.StatusBadRequest, ctxErr)
		return
	}

	jwt := lh.authCore.NewJwt(&model.User{
		ID: 1,
	})

	c.JSON(200, jwt)
}
