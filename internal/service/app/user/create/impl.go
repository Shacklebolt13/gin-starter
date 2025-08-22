package create

import (
	"context"
	"gin-starter/internal/database/sql/model"
	"gin-starter/internal/database/sql/repo"
	"gin-starter/internal/service/integration/amazon/cognito"
	"gin-starter/internal/service/integration/amazon/cognito/usecase/user/create"
)

type createUserServiceImpl struct {
	cognito cognito.CognitoService
	usrRepo repo.UserRepository
}

func (svc *createUserServiceImpl) RegisterUserByPassword(ctx context.Context, req CreateUserByPasswordRequest) (*CreateUserByPasswordResult, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := svc.cognito.Create.RegisterUserWithPassword(ctx, create.CreateUserRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	user, err := svc.usrRepo.Create(ctx, model.User{
		Email: req.Email,
		Name:  req.Username,
		BaseModelStringPk: model.BaseModelStringPk{
			ID: res.UserSub,
		},
	})
	if err != nil {
		return nil, err
	}

	return &CreateUserByPasswordResult{
		UserID: user.ID,
		Email:  user.Email,
		Name:   user.Name,
	}, nil
}

func NewCreateUserService(cognito *cognito.CognitoService, usrRepo repo.UserRepository) (CreateUserService, error) {
	return &createUserServiceImpl{
		cognito: *cognito,
		usrRepo: usrRepo,
	}, nil
}
