package user

import (
	"gin-starter/internal/service/app/user/create"
	"gin-starter/internal/service/app/user/login"
)

type UserService struct {
	Create create.CreateUserService
	Login  login.LoginUserService
}

func NewUserService(create create.CreateUserService, login login.LoginUserService) *UserService {
	return &UserService{
		Create: create,
		Login:  login,
	}
}
