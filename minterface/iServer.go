package minterface

import "net"

type IServer interface {
	Init()

	Start()

	Stop()

	LoggerClose()

	ConnectHandle(conn *net.TCPConn) (err error)

	ExitHandle()

	AddRouter(apiId int32, handle HandleFunc)
}
