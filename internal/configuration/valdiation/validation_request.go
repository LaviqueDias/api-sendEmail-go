package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/LaviqueDias/api-sendEmail-go/internal/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init(){
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)

		_ = val.RegisterValidation("notblank", func(fl validator.FieldLevel) bool {
			s, ok := fl.Field().Interface().(string)
			if !ok {
				return false
			}
			return strings.TrimSpace(s) != ""
		})

		_ = val.RegisterValidation("mintrim", func(fl validator.FieldLevel) bool {
			p := fl.Param() // número passado na tag
			s, ok := fl.Field().Interface().(string)
			if !ok {
				return false
			}
			trimmed := strings.TrimSpace(s)
			// converter p pra int
			var n int
			_, _ = fmt.Sscanf(p, "%d", &n)
			return len([]rune(trimmed)) >= n
		})

	}

}

func ValidateRequestError(validation_err error) *rest_err.RestErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr){
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors){
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field: e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
	
} 