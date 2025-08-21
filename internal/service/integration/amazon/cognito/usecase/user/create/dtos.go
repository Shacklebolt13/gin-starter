package create

import "gin-starter/internal/utils/validation"

type CreateUserRequest struct {
	validation.CanValidate
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}
