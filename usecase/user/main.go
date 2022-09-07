package user

import "github.com/htoyoda18/sample-tweet-api/domain/model"

type UserUseCase interface {
	Add() (*model.User, error)
}

type userUseCase struct {
}
