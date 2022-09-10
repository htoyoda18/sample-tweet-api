package context

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/repository"
)

type ContextApi interface {
	Context(c *gin.Context) (*gorm.DB, error)
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
	log.Printf("AddUser")

	db := c.MustGet("db").(*gorm.DB)

	return db, nil
}
