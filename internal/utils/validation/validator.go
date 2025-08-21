package validation

import "github.com/go-playground/validator/v10"

type Validatable interface {
	Validate() error
}

type CanValidate struct {
}

func (v *CanValidate) Validate() error {
	validatorObj := validator.New()
	err := validatorObj.Struct(v)
	return err
}
