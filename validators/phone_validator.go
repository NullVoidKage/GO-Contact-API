// validators/phone_validator.go
package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/nyaruka/phonenumbers"
)

func RegisterPhoneValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("uniquePhoneNumbers", uniquePhoneNumbers)
	}
}

func uniquePhoneNumbers(fl validator.FieldLevel) bool {
	numbers := fl.Field().Interface().([]string)
	seen := map[string]bool{}
	for _, number := range numbers {
		parsed, err := phonenumbers.Parse(number, "AU")
		if err != nil || !phonenumbers.IsValidNumber(parsed) {
			return false
		}
		e164 := phonenumbers.Format(parsed, phonenumbers.E164)
		if seen[e164] {
			return false
		}
		seen[e164] = true
	}
	return true
}
