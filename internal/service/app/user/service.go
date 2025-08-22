package user

import (
	"gin-starter/internal/service/app/user/create"
)

type UserService struct {
	Create create.CreateUserService
}

func NewUserService(create create.CreateUserService) (*UserService, error) {
	return &UserService{
		Create: create,
	}, nil
}
