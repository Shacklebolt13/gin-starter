package create

type User struct {
	Id    string `json:"id" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
