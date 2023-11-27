package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationError(err error) ([]string, []string) {
	var tag, field, param string
	var fields, messages []string
	for _, v := range err.(validator.ValidationErrors) {
		tag = v.Tag()
		field = v.Field()
		field = strings.ToLower(field[:1]) + field[1:]
		param = v.Param()
		valueType := reflect.TypeOf(field).Kind()

		switch tag {
		case "required":
			fields = append(fields, field)
			messages = append(messages, field+" is required")
		case "max":
			fields = append(fields, field)
			messages = append(messages, field+" cannot be longer than "+param)
		case "min":
			if valueType == reflect.String {
				fields = append(fields, field)
				messages = append(messages, field+" must be at least "+param+" character in length")

			} else if valueType == reflect.Int {
				fields = append(fields, field)
				messages = append(messages, field+" must be at least "+param+" digit in length")
			}
		case "email":
			fields = append(fields, field)
			messages = append(messages, "invalid email format")
		case "len":
			fields = append(fields, field)
			messages = append(messages, field+" must be "+param+" characters long")
		case "required_if":
			fields = append(fields, field)
			messages = append(messages, field+" is required if "+param)
		case "lt":
			fields = append(fields, field)
			messages = append(messages, field+" must be less than "+param+" in length ")
		case "lte":
			fields = append(fields, field)
			messages = append(messages, field+" must be at maximum "+param+" in length ")
		case "gt":
			fields = append(fields, field)
			messages = append(messages, field+" must be greater than "+param+" in length ")
		case "gte":
			fields = append(fields, field)
			messages = append(messages, field+" must be at least "+param+" in length ")
		case "eqfield":
			fields = append(fields, field)
			messages = append(messages, field+" must be equal to "+param)
		case "alpha":
			fields = append(fields, field)
			messages = append(messages, field+" can only contain alphabetic characters")
		case "alphanum":
			fields = append(fields, field)
			messages = append(messages, field+" can only contain alphanumeric characters")
		case "numeric":
			fields = append(fields, field)
			messages = append(messages, field+" must be a valid numeric value")
		case "number":
			fields = append(fields, field)
			messages = append(messages, field+" must must be a valid number")
		case "url":
			fields = append(fields, field)
			messages = append(messages, field+" must be a valid URL")
		case "uri":
			fields = append(fields, field)
			messages = append(messages, field+" must be a valid URI")
		case "contains":
			fields = append(fields, field)
			messages = append(messages, field+" must contain the text "+"'"+param+"'")
		case "excludes":
			fields = append(fields, field)
			messages = append(messages, field+" cannot contain the text "+"'"+param+"'")
		case "hexadecimal":
			fields = append(fields, field)
			messages = append(messages, field+" must be a valid hexadecimal")
		case "uuid":
			fields = append(fields, field)
			messages = append(messages, field+" must be a valid UUID")
		case "ulid":
			fields = append(fields, field)
			messages = append(messages, field+" must be a valid ULID")
		case "multibyte":
			fields = append(fields, field)
			messages = append(messages, field+" must contain multibyte characters")
		case "iscolor":
			fields = append(fields, field)
			messages = append(messages, field+"  must be a valid color")
		case "oneof":
			fields = append(fields, field)
			messages = append(messages, field+"  must be one of "+"["+param+"]")
		case "json":
			fields = append(fields, field)
			messages = append(messages, field+"  must be a valid json string")
		case "lowercase":
			fields = append(fields, field)
			messages = append(messages, field+" must be a lowercase string")
		case "uppercase":
			fields = append(fields, field)
			messages = append(messages, field+" must be an uppercase string")
		case "eq":
			fields = append(fields, field)
			messages = append(messages, field+" is not equal to "+param)
		case "datetime":
			fields = append(fields, field)
			messages = append(messages, field+" does not match the "+param+"format")
		case "ne":
			fields = append(fields, field)
			messages = append(messages, field+" should not be equal to "+param)
		case "gtfield":
			fields = append(fields, field)
			messages = append(messages, field+" must be greater than "+param)
		case "unique":
			fields = append(fields, field)
			messages = append(messages, field+" must contain unique values")
		case "gtefield":
			fields = append(fields, field)
			messages = append(messages, field+" must be greater than or equal to "+param)
		case "ltfield":
			fields = append(fields, field)
			messages = append(messages, field+" must be less than "+param)
		case "ltefield":
			fields = append(fields, field)
			messages = append(messages, field+"  must be less than or equal to "+param)
		case "necsfield":
			fields = append(fields, field)
			messages = append(messages, field+" cannot be equal to "+param)
		case "eqcsfield":
			fields = append(fields, field)
			messages = append(messages, field+" must be equal to "+param)
		}
	}

	return fields, messages
}

func ValidationQueryParamsError(err error) string {
	var tag, field, param string

	for _, v := range err.(validator.ValidationErrors) {
		tag = v.Tag()
		field = v.Field()
		param = v.Param()
		valueType := reflect.TypeOf(field).Kind()
		fmt.Println("field", field)

		switch tag {
		case "required":
			return field + " is required"
		case "max":
			return field + " cannot be longer than " + param
		case "min":
			if valueType == reflect.String {
				return field + " must be at least " + param + " character in length"
			} else if valueType == reflect.Int {
				return field + " must be at least " + param + " digit in length"
			}

		case "email":
			return "invalid email format"
		case "len":
			return field + " must be " + param + " characters long"
		case "required_if":
			return field + " is required if " + param
		case "lt":
			return field + " must be less than " + param + " in length "
		case "lte":
			return field + " must be at maximum " + param + " in length "
		case "gt":
			return field + " must be greater than " + param + " in length "
		case "gte":
			return field + " must be at least " + param + " in length "
		case "eqfield":
			return field + " must be equal to " + param
		case "alpha":
			return field + " can only contain alphabetic characters"
		case "alphanum":
			return field + " can only contain alphanumeric characters"
		case "numeric":
			return field + " must be a valid numeric value"
		case "number":
			return field + " must must be a valid number"
		case "url":
			return field + " must be a valid URL"
		case "uri":
			return field + " must be a valid URI"
		case "contains":
			return field + " must contain the text " + "'" + param + "'"
		case "excludes":
			return field + " cannot contain the text " + "'" + param + "'"
		case "hexadecimal":
			return field + " must be a valid hexadecimal"
		case "uuid":
			return field + " must be a valid UUID"
		case "ulid":
			return field + " must be a valid ULID"
		case "multibyte":
			return field + " must contain multibyte characters"
		case "iscolor":
			return field + "  must be a valid color"
		case "oneof":
			return field + "  must be one of " + "[" + param + "]"
		case "json":
			return field + "  must be a valid json string"
		case "lowercase":
			return field + " must be a lowercase string"
		case "uppercase":
			return field + " must be an uppercase string"
		case "eq":
			return field + " is not equal to " + param
		case "datetime":
			return field + " does not match the " + param + "format"
		case "ne":
			return field + " should not be equal to " + param
		case "gtfield":
			return field + " must be greater than " + param
		case "unique":
			return field + " must contain unique values"
		case "gtefield":
			return field + " must be greater than or equal to " + param
		case "ltfield":
			return field + " must be less than " + param
		case "ltefield":
			return field + "  must be less than or equal to " + param
		case "necsfield":
			return field + " cannot be equal to " + param
		case "eqcsfield":
			return field + " must be equal to " + param
		}
	}

	return err.Error()
}
