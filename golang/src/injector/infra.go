package injector

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/repository"
	"github.com/htoyoda18/sample-tweet-api/golang/src/infra/persistence/user"
)

type Infra struct {
	UserPersistence repository.UserRepository
}

func NewInfra() *Infra {
	UserPersistence := user.NewUserPersistence()

	return &Infra{
		UserPersistence,
	}
}
