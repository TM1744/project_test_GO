package validator

import (
	"fmt"
	"projeto-teste/fail"
	"strings"
)

func ValidStrintIsNotBlank(value *string, fieldName string) error {
	if strings.TrimSpace(*value) == "" {
		return fail.StringToBadRequest(fieldName + " cannot be empty")
	}
	return nil
}

func ValidStringMaxLength(value *string, maxLength int, fieldName string) error {
	if len(*value) > maxLength {
		return fail.StringToBadRequest(fieldName + " cannot exceed " + fmt.Sprint(maxLength) + " characters")
	}

	return nil
}
