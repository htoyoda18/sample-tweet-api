package injector

import "github.com/htoyoda18/sample-tweet-api/service/context"

type Service struct {
	ApiContext context.ContextApi
}

func NewService() *Service {
	return &Service{}
}
