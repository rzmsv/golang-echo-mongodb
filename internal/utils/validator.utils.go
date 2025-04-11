package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type requestValoidator struct {
	validator *validator.Validate
}

type validatorUtil struct {
	EValidator *echo.Validator
}

func Validator(eValidator *echo.Validator) *validatorUtil {
	return &validatorUtil{
		EValidator: eValidator,
	}
}
func (reqValidator *requestValoidator) Validate(i interface{}) error {
	return reqValidator.validator.Struct(i)
}

func (echoValidator *validatorUtil) InitialValidator() {
	*echoValidator.EValidator = &requestValoidator{validator: validator.New()}
}
