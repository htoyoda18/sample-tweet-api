package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/user/request"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/context"
	"github.com/htoyoda18/sample-tweet-api/golang/src/usecase/user"
)

type UserHandler interface {
	AddUser(*gin.Context)
	DeleteUser(*gin.Context)
	UpdateUser(*gin.Context)
	ShowUser(*gin.Context)
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

// AddUser ユーザ作成のAPI
// @Summary ユーザ作成のAPI
// @description ユーザ作成のAPI
// @param request body request.AddUsersReq true "request body"
// @Success 200
// @Failure 400
// @Failure 500
// @router /signup [POST]
func (uh userHandler) AddUser(c *gin.Context) {
	log.Printf("AddUser")

	var params request.AddUsersReq

	if err := c.ShouldBindJSON(&params); err != nil {
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

// AddUser ユーザ削除のAPI
// @Summary ユーザ削除のAPI
// @description ユーザ削除のAPI
// @param id path int true "userID"
// @Success 200
// @Failure 400
// @Failure 500
// @router /user/:id [DELETE]
func (uh userHandler) DeleteUser(c *gin.Context) {
	log.Printf("DeleteUser")

	id, errID := strconv.ParseUint(c.Param("id"), 10, 32)

	if errID != nil {
		log.Println("Erorr DeleteUser: ", errID)
		c.AbortWithStatusJSON(http.StatusBadRequest, errID)
		return
	}

	ctx, ctxErr := uh.ctx.ContextUser(c)

	if ctxErr != nil {
		log.Printf("Erorr DeleteUser")
		c.AbortWithError(http.StatusBadRequest, ctxErr)
		return
	}

	err := uh.userUseCase.DeleteUser(ctx, model.UserId(id))
	if err != nil {
		log.Printf("Erorr DeleteUser")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Status(200)
}

// AddUser ユーザ更新のAPI
// @Summary ユーザ更新のAPI
// @description ユーザ更新のAPI
// @param id path int true "userID"
// @Success 200
// @Failure 400
// @Failure 500
// @router /user/:id [PATCH]
func (uh userHandler) UpdateUser(c *gin.Context) {
	log.Printf("UpdateUser")

	var params request.UpdateUserReq
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("Erorr AddUser")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, errID := strconv.ParseUint(c.Param("id"), 10, 32)

	if errID != nil {
		log.Println("Erorr DeleteUser: ", errID)
		c.AbortWithStatusJSON(http.StatusBadRequest, errID)
		return
	}

	ctx, ctxErr := uh.ctx.ContextUser(c)

	if ctxErr != nil {
		log.Printf("Erorr UpdateUser")
		c.AbortWithError(http.StatusBadRequest, ctxErr)
		return
	}

	// ユーザーを更新
	user, err := uh.userUseCase.UpdateUser(
		ctx,
		model.UserId(id),
		params,
	)

	if err != nil {
		log.Printf("Erorr UpdateUser")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(200, user)
}

// AddUser ユーザ閲覧のAPI
// @Summary ユーザ閲覧のAPI
// @description ユーザ閲覧のAPI
// @param id path int true "userID"
// @Success 200
// @Failure 400
// @Failure 500
// @router /user/:id [GET]
func (uh userHandler) ShowUser(c *gin.Context) {
	log.Printf("ShowUser")

	id, errID := strconv.ParseUint(c.Param("id"), 10, 32)

	if errID != nil {
		log.Println("Erorr ShowUser: ", errID)
		c.AbortWithStatusJSON(http.StatusBadRequest, errID)
		return
	}

	ctx, ctxErr := uh.ctx.ContextUser(c)

	if ctxErr != nil {
		log.Printf("Erorr ShowUser")
		c.AbortWithError(http.StatusBadRequest, ctxErr)
		return
	}

	user, err := uh.userUseCase.ShowUser(
		ctx.DB,
		&model.User{
			ID: model.UserId(id),
		},
	)

	if err != nil {
		log.Println("Erorr ShowUser: ", errID)
		c.AbortWithStatusJSON(http.StatusBadRequest, errID)
		return
	}

	c.JSON(200, user)
}
