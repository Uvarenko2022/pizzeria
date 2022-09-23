package validate

import (
	"log"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type CustomValidator struct {
	v     *validator.Validate
	trans ut.Translator
}

type Translate interface {
	TranslateError(err error) string
	Struct(v interface{}) error
}

func New(v *validator.Validate, trans ut.Translator) Translate {
	validator := &CustomValidator{v, trans}

	return validator
}

func addTranslation(tag string, errMessage string, v *validator.Validate, trans ut.Translator) {
	registerFn := func(ut ut.Translator) error {
		return ut.Add(tag, errMessage, false)
	}

	transFn := func(ut ut.Translator, fe validator.FieldError) string {
		param := fe.Param()
		tag := fe.Tag()

		t, err := ut.T(tag, fe.Field(), param)
		if err != nil {
			return fe.(error).Error()
		}
		return t
	}

	v.RegisterTranslation(tag, trans, registerFn, transFn)
}

func (cv *CustomValidator) Struct(v interface{}) error {
	return cv.v.Struct(v)
}

func (cv *CustomValidator) TranslateError(err error) string {
	if err == nil {
		return ""
	}

	var errs string
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		errs += e.Translate(cv.trans)
	}
	return errs
}

func RegisterMessages(v *validator.Validate, trans ut.Translator) {
	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		log.Fatal(err)
	}

	//food
	addTranslation("name", "bad name, name should be more than 5 characters long and less than 30", v, trans)
	addTranslation("cost", "bad cost, cost should be more than 15", v, trans)
	addTranslation("type", "bad type, type should be 0 or 1", v, trans)

	//order
	addTranslation("state", "bad state, state should be 0, 1 or 2", v, trans)
	addTranslation("foodids", "bad foodids, foodids should be integers", v, trans)
}

func RegisterValidations(v *validator.Validate) {
	v.RegisterValidation("name", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 5 && len(fl.Field().String()) < 30
	})
	v.RegisterValidation("cost", func(fl validator.FieldLevel) bool {
		return fl.Field().Float() > 15
	})
	v.RegisterValidation("type", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() == 1 || fl.Field().Int() == 0
	})
	v.RegisterValidation("state", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() == 0 || fl.Field().Int() == 1 || fl.Field().Int() == 2
	})
	v.RegisterValidation("foodids", func(fl validator.FieldLevel) bool {
		ids := fl.Field().Interface().([]uint)
		for _, v := range ids {
			if v == 0 {
				return false
			}
		}
		return true
	})
}
