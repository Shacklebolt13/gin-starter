package create

import "context"

type CreateUserService interface {
	RegisterUserByPassword(ctx context.Context, req CreateUserByPasswordRequest) (*CreateUserByPasswordResult, error)
}
