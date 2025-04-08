package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegexValidator(expr string) validator.Func {
	re := regexp.MustCompile(expr)

	return func(fl validator.FieldLevel) bool {
		val := fl.Field().String()

		return re.MatchString(val)
	}
}
