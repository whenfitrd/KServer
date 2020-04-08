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

var StatusOk = RInfo{Ok, "Ok"}
var StatusError = RInfo{Error, "Error"}
var ApiAuthError = RInfo{Ok, "Api Authorization Error."}