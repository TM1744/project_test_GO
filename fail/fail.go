package fail

import "errors"

type BadRequest error
type InternalServerError error


func StringToBadRequest (message string) BadRequest {
	return errors.New(message)
}

func ErrorToBadRequest (err error) BadRequest {
	return err
}

func StringToInternalServerError (message string) InternalServerError {
	return errors.New(message)
}

func ErrorToInternalServerError (err error) InternalServerError {
	return err
}