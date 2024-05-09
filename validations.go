package utils

import (
	"github.com/go-playground/validator/v10"
	"gitlab.landui.cn/gomod/global"
)

func ParameterValidations(req interface{}) validator.FieldError {
	err := global.Validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}
	return nil
}
