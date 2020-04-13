package utils

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"os/exec"
)


func HandlePanic(info string) (err error) {
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

func CheckPortIsUsed(port int) bool {
	checkStatement := fmt.Sprintf("lsof -i:%d ", port)
	output, _ := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	if len(output) > 0 {
		return true
	}
	return false
}
