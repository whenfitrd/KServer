package utils

import "errors"

func HandlePanic() (err error) {
	if r := recover(); r != nil {
		logger.Error("Recovered in Panic2Error")

		switch x := r.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = errors.New("Unknow panic")
		}
	}
	return nil
}
