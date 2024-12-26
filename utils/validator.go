package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()

	// _ = Validate.RegisterValidation("unique_email", func(fl validator.FieldLevel) bool {
	// 	return false
	// })

	// _ = Validate.RegisterValidation("unique_username", func(fl validator.FieldLevel) bool {
	// 	return false
	// })

	// Validator for non-zero value like price
	_ = Validate.RegisterValidation("non_zero", func(fl validator.FieldLevel) bool {
		value := fl.Field().Float()
		return value > 0
	})
}

// ErrorFormatter
func ErrorFormatter(err error) map[string]string {
	errors := make(map[string]string)

	if ValidationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range ValidationErrors {
			errors[fieldError.Field()] = fmt.Sprintf("failed on the %s validation", fieldError.Tag())
		}
	}

	return errors
}
