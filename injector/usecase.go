package injector

import "github.com/htoyoda18/sample-tweet-api/usecase/user"

type UseCase struct {
	UserUseCase user.UserUseCase
}

func NewUseCase() *UseCase {
	infra := NewInfra()

	UserUseCase := user.NewUserUseCase(
		infra.UserPersistence,
	)
	return &UseCase{
		UserUseCase,
	}
}
