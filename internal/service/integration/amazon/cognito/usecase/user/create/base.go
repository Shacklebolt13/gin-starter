package create

import "context"

type CreateUserService interface {
	RegisterUserWithPassword(ctx context.Context, user CreateUserRequest) (*CreateUserResult, error)
}
