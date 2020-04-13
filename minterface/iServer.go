package minterface

import (
	"github.com/whenfitrd/KServer/rStatus"
	"net"
)

type IServer interface {
	SConfig(name, ip, port string)

	Init()

	Start()

	Stop()

	AcceptConnect()

	ConnectHandle(conn *net.TCPConn) (err error)

	AddRouter(apiId int32, handle *Function)

	WriteToGroup(data []byte, groupName string) rStatus.RInfo
}
