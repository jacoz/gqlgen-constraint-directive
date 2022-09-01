package gqlgen_constraint_directive

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type constraintDirective struct {
	validator  *validator.Validate
	translator ut.Translator
}

func newConstraintDirective(validator *validator.Validate, translator ut.Translator) *constraintDirective {
	en_translations.RegisterDefaultTranslations(validator, translator)

	return &constraintDirective{
		validator,
		translator,
	}
}

func (b *constraintDirective) Constraint(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string) (interface{}, error) {
	val, err := next(ctx)
	if err != nil {
		panic(err)
	}

	field := *graphql.GetPathContext(ctx).Field

	err = b.validator.Var(val, constraint)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		transErr := fmt.Errorf("%s%+v", field, validationErrors[0].Translate(b.translator))
		return val, transErr
	}

	return val, nil
}

func (b *constraintDirective) AddTranslationForTag(tag string, message string) error {
	return b.validator.RegisterTranslation(
		tag,
		b.translator,
		func(ut ut.Translator) error {
			return ut.Add(tag, message, true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(tag, fe.Field())
			return t
		},
	)
}
