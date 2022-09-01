package gqlgen_constraint_directive

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func New(v *validator.Validate, t *ut.Translator) *constraintDirective {
	var govalidator *validator.Validate
	if v == nil {
		govalidator = validator.New()
	} else {
		govalidator = v
	}

	var translator ut.Translator
	if t == nil {
		enlang := en.New()
		uni := ut.New(enlang, enlang)
		translator, _ = uni.GetTranslator("en")
	} else {
		translator = *t
	}

	return newConstraintDirective(
		govalidator,
		translator,
	)
}
