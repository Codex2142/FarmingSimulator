package validator

import "github.com/go-playground/validator/v10"

var Validate = validator.New()

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Field()+" is "+e.Tag())
	}
	return errors
}
