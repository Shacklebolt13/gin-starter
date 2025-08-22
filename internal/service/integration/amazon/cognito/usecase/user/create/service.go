package create

import (
	"context"
	"gin-starter/internal/singleton/config"
	"gin-starter/internal/utils/log"

	"github.com/aws/aws-sdk-go-v2/aws"
	cidp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type createUserServiceImpl struct {
	client *cidp.Client
	cfg    *config.AppConfig
	aws    *aws.Config
}

func (svc *createUserServiceImpl) RegisterUserWithPassword(ctx context.Context, request CreateUserRequest) (*CreateUserResult, error) {
	if err := request.Validate(); err != nil {
		log.Error().Err(err).Msg("validation failed for Registering User By password on cognito")
		return nil, err
	}

	signUpInput := cidp.SignUpInput{
		ClientId: aws.String(svc.cfg.Env.COGNITO_CLIENT_ID),
		Username: aws.String(request.Email),
		Password: aws.String(request.Password),
	}

	out, err := svc.client.SignUp(ctx, &signUpInput)
	if err != nil {
		log.Error().Err(err).Msg("failed to sign up user")
		return nil, err
	}

	return &CreateUserResult{
		UserSub: *out.UserSub,
	}, nil
}

func NewCreateUserService(client *cidp.Client, cfg *config.AppConfig, awsCfg *aws.Config) CreateUserService {
	return &createUserServiceImpl{
		client: client,
		cfg:    cfg,
		aws:    awsCfg,
	}
}
