package login

import "context"

type LoginUserService interface {
	LoginByPassword(ctx context.Context, request LoginByPasswordRequest) (*LoginByPasswordResult, error)
}
