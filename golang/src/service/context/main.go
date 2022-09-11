package context

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/repository"
)

type ContextUser struct {
	DB     *gorm.DB
	UserId model.UserId
}

type ContextApi interface {
	Context(c *gin.Context) (*gorm.DB, error)
	ContextUser(c *gin.Context) (*ContextUser, error)
}

type contextApi struct {
	userRepository repository.UserRepository
}

func NewContextApi(
	userRepository repository.UserRepository,
) ContextApi {
	return &contextApi{
		userRepository,
	}
}

func (contextApi) Context(c *gin.Context) (*gorm.DB, error) {
	log.Println("Context")

	db := c.MustGet("db").(*gorm.DB)

	return db, nil
}

func (contextApi) ContextUser(c *gin.Context) (*ContextUser, error) {
	log.Println("ContextUser")

	userid := c.MustGet("user_id").(int)
	db := c.MustGet("db").(*gorm.DB)

	return &ContextUser{
		DB:     db,
		UserId: model.UserId(userid),
	}, nil
}
