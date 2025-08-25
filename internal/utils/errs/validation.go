package errs

import (
	"github.com/go-playground/validator/v10"
)

type ValidationErr validator.ValidationErrors

func (v ValidationErr) Error() string {
	return "Validation Failure"
}

func (v *ValidationErr) Map() map[string]string {
	var errMap = make(map[string]string)
	for _, e := range *v {
		errMap[e.Field()] = e.Tag()
	}

	return errMap
}
