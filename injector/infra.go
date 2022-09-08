package injector

import (
	"github.com/htoyoda18/sample-tweet-api/domain/repository"
	"github.com/htoyoda18/sample-tweet-api/infra/persistence/user"
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
