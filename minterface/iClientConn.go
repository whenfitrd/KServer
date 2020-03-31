package minterface

import "net"

type ICConn interface {
	//手动初始化
	Init(tc *net.TCPConn)
	//写
	Write(data []byte)
	//读
	Read()
	//处理数据
	Handle()
	//关闭链接
	Close()

	GetUID() string
}
