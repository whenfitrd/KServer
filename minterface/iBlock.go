package minterface

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IBlock interface {
	//激活
	Active()
	//冻结
	Freeze()
	//获取块对象
	GetBlock() IBlock
}

type INetGroupBlock interface {
	IBlock
	//初始化配置
	Config()
	//获取网络组管理器
	GetNetGroupManager() INetGroupManager
}

type IServerBlock interface {
	IBlock
	//启动
	Start(name, ip, port string)
	//关闭
	Stop()
	//获取模块服务器
	GetServer() IServer
}

type IDBBlock interface {
	IBlock
	//链接
	Connect() *mongo.Client
	//
}
