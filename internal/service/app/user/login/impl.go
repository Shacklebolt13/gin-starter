package login

import (
	"context"
	"gin-starter/internal/database/sql/repo"
	"gin-starter/internal/service/integration/amazon/cognito"
	"gin-starter/internal/service/integration/amazon/cognito/usecase/user/login"
	"gin-starter/internal/utils/errs"

	"github.com/go-playground/validator/v10"
)

type loginUserServiceImpl struct {
	cognito cognito.CognitoService
	usrRepo repo.UserRepository
}

func (l *loginUserServiceImpl) LoginByPassword(ctx context.Context, request LoginByPasswordRequest) (*LoginByPasswordResult, error) {
	//Validate Request
	if err := request.Validate(); err != nil {
		return nil, &errs.Incident{
			Err:     err,
			Type:    errs.RequestValidationError,
			Message: err.Error(),
			Detail:  err.(validator.ValidationErrors),
		}
	}

	// Call Cognito to authenticate the user
	authResult, err := l.cognito.Login.LoginByPassword(ctx, login.LoginByPasswordRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, &errs.Incident{
			Err:     err,
			Type:    errs.AuthenticationError,
			Message: err.Error(),
		}
	}

	// Map the Cognito authentication result to the login result
	result := &LoginByPasswordResult{
		AccessToken:  authResult.AccessToken,
		IdToken:      authResult.IdToken,
		RefreshToken: authResult.RefreshToken,
	}

	return result, nil
}

func NewLoginUserService(cognito *cognito.CognitoService, usrRepo repo.UserRepository) LoginUserService {
	return &loginUserServiceImpl{
		cognito: *cognito,
		usrRepo: usrRepo,
	}
}
