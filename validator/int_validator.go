package validator

import (
	"projeto-teste/fail"
)

func ValidIntHighThanZero(value *int64, fieldName string) error {
	if value == nil || *value <= 0 {
		return fail.StringToBadRequest(fieldName + " must be greater than zero and not null")
	}
	return nil
}

func ValidIntHighThanZeroButCanNil(value *int64, fieldName string) error {
	if value == nil { return nil}
	if *value <= 0 {
		return fail.StringToBadRequest(fieldName + " must be greater than zero")
	}
	return nil
}

func ValidIntNotLessThanZero(value *int64, fieldName string) error {
	if value == nil || *value < 0 {
		return fail.StringToBadRequest(fieldName + " cannot be less than zero and not null")
	}

	return nil
}

func ValidIntNotLessThanZeroButCanNil(value *int64, fieldName string) error {
	if value == nil { return nil }
	if *value < 0 {
		return fail.StringToBadRequest(fieldName + " cannot be less than zero")
	}

	return nil
}

func ValidIntEqualToZero(value *int64, fieldName string) error {
	if value == nil || *value != 0 {
		return fail.StringToBadRequest(fieldName + " must be equal to zero")
	}
	return nil
}

func ValidIntIsHighestThanValue(highestValue *int64, lowestValue *int64, highestFieldName string, lowestFieldName string) error {
	if (highestValue == nil && lowestValue != nil) || (*highestValue < *lowestValue) {
		return fail.StringToBadRequest(highestFieldName + " cannot be less than the " + lowestFieldName)
	}
	return nil
}
