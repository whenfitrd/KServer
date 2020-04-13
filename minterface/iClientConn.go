package minterface

import "net"

type ICConn interface {
	//手动初始化
	Init(tc *net.TCPConn)
	//写
	Write(data []byte)
	//读
	Read(router IRouter)
	//处理数据
	Handle(router IRouter)
	//关闭链接
	Close()
	//获取UUID
	GetUID() string
	//更新权限
	UpdateAuth(auth int)
	//获取权限
	GetAuth() int
}
