package injector

import (
	"github.com/htoyoda18/sample-tweet-api/golang/src/service/context"
)

type Service struct {
	ApiContext context.ContextApi
}

func NewService() *Service {
	infra := NewInfra()

	apiContext := context.NewContextApi(
		infra.UserPersistence,
	)
	return &Service{
		apiContext,
	}
}
