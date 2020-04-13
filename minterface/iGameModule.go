package minterface

type IGameModule interface {
	//启动
	Start(name, ip, port string)
	//关闭
	Stop()
	//获取模块服务器
	GetServer() IServer
}
