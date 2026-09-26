package converter

import (
	"errors"
	"strconv"
)

type Converter struct {}



func StringToInt64 (value string) (number int64, err error) {
	number, err = strconv.ParseInt(value, 10, 64)

	if err != nil  {
		return 0, errors.New("failed to convert string to int64")
	}

	return
}

func Int64ToString (value int64) (text string) {
	return strconv.FormatInt(value, 10)
}