package minterface

import "github.com/whenfitrd/KServer/rStatus"

type HandleFunc func(ICConn, []byte)

type Function struct {
	Func HandleFunc
	Auth int
}

type IRouter interface {
	//设置即将注册的路由权限
	SetAuth(auth int)
	//注册路由
	AddRouter(apiId int32, handle HandleFunc)
	//路由处理
	Handle(cc ICConn, data []byte) rStatus.RInfo

	GetHandleMap() map[int]*Function
}
