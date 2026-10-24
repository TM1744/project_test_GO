package validator

import (
	"fmt"
	"projeto-teste/fail"
	"strings"
)

func ValidStrintIsNotBlank(value *string, fieldName string) error {
	if value == nil || strings.TrimSpace(*value) == "" {
		return fail.StringToBadRequest(fieldName + " cannot be empty")
	}
	return nil
}

func ValidStringMaxLength(value *string, maxLength int, fieldName string) error {
	if value == nil { return nil }
	if len(*value) > maxLength {
		return fail.StringToBadRequest(fieldName + " cannot exceed " + fmt.Sprint(maxLength) + " characters")
	}

	return nil
}

func ValidStringMinLength(value *string, minLegth int, fieldName string) error {
	if value == nil || strings.TrimSpace(*value) == "" { return nil }
	if len(*value) < minLegth {
		return fail.StringToBadRequest(fieldName + " cannot have fewer than " + fmt.Sprint(minLegth) + " characters")
	}

	return nil
}
