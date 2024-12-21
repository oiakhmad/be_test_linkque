package validator

import (
	"regexp"
	"strings"
	"sync"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	lock       = &sync.Mutex{}
	validate   *validator.Validate
	translator ut.Translator
)

//GetValidator Initiatilize validator in singleton way
func GetValidator() *validator.Validate {
	lock.Lock()
	defer lock.Unlock()

	if validate != nil {
		return validate
	}

	// NOTE: ommitting allot of error checking for brevity
	en := en.New()
	uni := ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	translator, _ = uni.GetTranslator("en")
	validate = validator.New()

	// register english translation
	en_translations.RegisterDefaultTranslations(validate, translator)

	// register custom ISO8601 validation
	validate.RegisterValidation("ISO8601", func(fl validator.FieldLevel) bool {
		ISO8601RegexString := "^((?:(\\d{4}-\\d{2}-\\d{2})T(\\d{2}:\\d{2}:\\d{2}(?:\\.\\d+)?))(Z|[\\+-]\\d{2}:\\d{2})?)$"
		ISO8601Regex := regexp.MustCompile(ISO8601RegexString)
		return ISO8601Regex.MatchString(fl.Field().String())
	})

	// register custom ISO8601 translation
	validate.RegisterTranslation("ISO8601", translator, func(ut ut.Translator) error {
		return ut.Add("ISO8601", "{0} must following ISO8601 or RFC3339 date format", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("ISO8601", fe.Field())
		return t
	})

	return validate
}

func GetValidatorMessage(err error) (messages []string) {
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, err.Translate(translator))
		}
	}
	return
}

func GetGeneralValidatorMessage(err error) (messages map[string]interface{}) {
	if err != nil {
		messages := make(map[string]interface{})

		for _, err := range err.(validator.ValidationErrors) {
			messages[strings.ToLower(err.Field())] = err.Translate(translator)
		}
		return removeTopStruct(messages)
	}
	return
}

func removeTopStruct(fields map[string]interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
