package rStatus

type RStatus int

type RInfo struct {
	RSt RStatus
	Msg string
}

const (
	Ok = iota
	Error
)

var StatusOK = RInfo{Ok, "Ok"}
var StatusError = RInfo{Error, "Error"}