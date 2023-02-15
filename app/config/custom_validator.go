package config

import (
	"github.com/go-playground/validator/v10"
	"shipt/app/utils"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(value interface{}) error {
	if err := cv.Validator.Struct(value); err != nil {
		errorMap := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errorMap[err.Field()] = "Expected to be a valid " + err.Tag()
		}
		return utils.ValidationError(errorMap)
	}
	return nil
}
