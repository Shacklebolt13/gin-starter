package login

import "gin-starter/internal/utils/validation"

type LoginByPasswordRequest struct {
	validation.CanValidate
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginByPasswordResult struct {
	AccessToken  string `json:"access_token"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}
