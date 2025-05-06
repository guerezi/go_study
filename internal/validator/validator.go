package validator

import (
	"errors"
	"regexp"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/sirupsen/logrus"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()

	rules := map[string]string{
		"street":  `^[A-Za-zÀ-ÿ0-9\sºª.,'-]{3,100}$`,
		"number":  `^[0-9]{1,6}[A-Za-z\-ºª\/]{0,10}$`,
		"city":    `^[A-Za-zÀ-ÿ\s'-]{2,100}$`,
		"state":   `^[A-Z]{2}$`,
		"zipcode": `^[0-9]{5}-[0-9]{3}$`,
	}

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		logrus.WithError(err).Fatal("Error registering default translations")
	}

	for index, val := range rules {
		if err := v.RegisterValidation(index, RegexValidator(val)); err != nil {
			logrus.WithError(err).Fatalf("Error registering %s validator", index)
		}

		if err := v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
			return ut.Add("required", "{0} must have a value!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())

			return t
		}); err != nil {
			logrus.WithError(err).Fatalf("Error registering %s translation", index)
		}
	}

	return &Validator{
		validator: v,
	}
}

func RegexValidator(expr string) validator.Func {
	re := regexp.MustCompile(expr)

	return func(fl validator.FieldLevel) bool {
		val := fl.Field().String()

		return re.MatchString(val)
	}
}

var (
	ErrValidation    = errors.New("validation error")
	ErrInvalidStruct = errors.New("invalid struct error")
)

func (v *Validator) Validate(i any) error {
	if err := v.validator.Struct(i); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			logrus.WithError(err).Error("Invalid Struct Validation error")

			return ErrInvalidStruct
		}
		logrus.WithError(err).Error("Validation error")

		return ErrValidation
	}

	return nil
}
