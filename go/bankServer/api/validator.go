package api

import (
	"github.com/go-playground/validator/v10"
	"gitlab.com/d3vus/project/go/bankServer/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
