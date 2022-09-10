package injector

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/context"
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/core"
)

type Service struct {
	ApiContext context.ContextApi
	AuthCore   core.AuthCore
}

func NewService() *Service {
	infra := NewInfra()

	apiContext := context.NewContextApi(
		infra.UserPersistence,
	)

	authCore := core.NewAuthCore(
		infra.UserPersistence,
		infra.RefreshTokenPersistence,
	)

	return &Service{
		apiContext,
		authCore,
	}
}
