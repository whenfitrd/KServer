package minterface

import (
	"github.com/whenfitrd/KServer/rStatus"
	"net"
)

type IServer interface {
	Init()

	Start()

	Stop()

	LoggerClose()

	ConnectHandle(conn *net.TCPConn) (err error)

	ExitHandle()

	SetAuth(auth int)

	AddRouter(apiId int32, handle HandleFunc)

	WriteToGroup(data []byte, groupName string) rStatus.RInfo
}
