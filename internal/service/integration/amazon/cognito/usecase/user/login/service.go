package login

import (
	"context"
	"gin-starter/internal/utils/log"

	cidp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type loginServiceImpl struct {
	client *cidp.Client
}

func (svc *loginServiceImpl) LoginByPassword(ctx context.Context, request LoginByPasswordRequest) (*LoginByPasswordResult, error) {
	authInput := cidp.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		AuthParameters: map[string]string{
			"USERNAME": request.Email,
			"PASSWORD": request.Password,
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

func NewLoginService(client *cidp.Client) LoginService {
	return &loginServiceImpl{
		client: client,
	}
}
