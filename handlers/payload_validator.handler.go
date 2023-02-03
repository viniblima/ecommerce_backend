package handlers

import (
	"github.com/go-playground/validator/v10"
)

func ValidatePayload(payload interface{}) []string {

	validate := validator.New()
	var errors []string

	if err := validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			errors = append(errors, validationError.Error())
		}
	}

	return errors
}
