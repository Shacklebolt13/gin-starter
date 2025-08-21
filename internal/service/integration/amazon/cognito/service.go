package cognito

import (
	"gin-starter/internal/service/integration/amazon/cognito/usecase/user/create"
	"gin-starter/internal/service/integration/amazon/cognito/usecase/user/login"
)

type CognitoService struct {
	Create create.CreateUserService
	Login  login.LoginService
}

func NewCognitoService(cr create.CreateUserService, lo login.LoginService) (*CognitoService, error) {
	return &CognitoService{
		Create: cr,
		Login:  lo,
	}, nil
}
