package validator

import goValidator "gopkg.in/go-playground/validator.v9"

var _validator = goValidator.New()

func Struct(argu interface{}) error {
	return _validator.Struct(argu)
}
