package login

import "context"

type LoginService interface {
	LoginByPassword(ctx context.Context, request LoginByPasswordRequest) (*LoginByPasswordResult, error)
}
