package login

import (
	"context"
	"gin-starter/internal/singleton/config"
	"gin-starter/internal/utils/crypto"
	"gin-starter/internal/utils/log"

	cidp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type loginServiceImpl struct {
	client *cidp.Client
	cfg    *config.AppConfig
}

func (svc *loginServiceImpl) LoginByPassword(ctx context.Context, request LoginByPasswordRequest) (*LoginByPasswordResult, error) {
	hash, err := crypto.SecretHash(
		request.Email,
		svc.cfg.Env.COGNITO_CLIENT_ID,
		svc.cfg.Env.COGNITO_CLIENT_SECRET)
	if err != nil {
		return nil, err
	}

	authInput := cidp.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		ClientId: &svc.cfg.Env.COGNITO_CLIENT_ID,

		AuthParameters: map[string]string{
			"USERNAME":    request.Email,
			"PASSWORD":    request.Password,
			"SECRET_HASH": hash,
		},
	}

	out, err := svc.client.InitiateAuth(ctx, &authInput)
	if err != nil {
		log.Error().Err(err).Msg("Failed to authenticate user")
		return nil, err
	}

	return &LoginByPasswordResult{
		AccessToken:  *out.AuthenticationResult.AccessToken,
		IdToken:      *out.AuthenticationResult.IdToken,
		RefreshToken: *out.AuthenticationResult.RefreshToken,
	}, nil
}

func NewLoginService(client *cidp.Client, cfg *config.AppConfig) LoginService {
	return &loginServiceImpl{
		client: client,
		cfg:    cfg,
	}
}
