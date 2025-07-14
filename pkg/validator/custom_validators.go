package validator

import (
	"fmt"

	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	validatorv10 "github.com/go-playground/validator/v10"
	esTranslations "github.com/go-playground/validator/v10/translations/es"
)

var (
	Validate *validatorv10.Validate
	trans    ut.Translator
)

func InitValidator() error {
	Validate = validatorv10.New()

	// Configurar traducción al español
	spanish := es.New()
	uni := ut.New(spanish, spanish)
	trans, _ = uni.GetTranslator("es")

	if err := esTranslations.RegisterDefaultTranslations(Validate, trans); err != nil {
		return fmt.Errorf("error registrando traducciones: %w", err)
	}

	// -- Puedes registrar aquí validaciones personalizadas si deseas
	// Sobrescribimos algunos mensajes predeterminados
	registerTranslation := func(tag, msg string) {
		_ = Validate.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
			return ut.Add(tag, msg, true)
		}, func(ut ut.Translator, fe validatorv10.FieldError) string {
			t, _ := ut.T(tag, fe.Field())
			return t
		})
	}

	// Mensajes personalizados
	registerTranslation("required", "{0} es obligatorio")
	registerTranslation("min", "{0} es demasiado corto")
	registerTranslation("max", "{0} es demasiado largo")
	registerTranslation("alphanum", "{0} solo debe contener caracteres alfanuméricos")

	return nil
}

// FormatValidationError traduce y limpia todos los mensajes de error
func FormatValidationError(err error) map[string]string {
	errors := map[string]string{}

	if validationErrors, ok := err.(validatorv10.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors[e.Field()] = e.Translate(trans)
		}
	} else {
		errors["error"] = err.Error()
	}

	return errors
}
