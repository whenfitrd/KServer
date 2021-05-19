package utils

import (
	"errors"
	"github.com/satori/go.uuid"
)


func HandlePanic(info string) (err error) {
	logger := GetLogger()
	if r := recover(); r != nil {
		logger.Info(info)
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

func UniqueString() string {
	uid := uuid.NewV4()
	return uid.String()
}
