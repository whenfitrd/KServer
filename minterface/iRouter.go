package minterface

import "github.com/whenfitrd/KServer/rStatus"

type HandleFunc func(ICConn, []byte)

type Function struct {
	Func HandleFunc
	Auth []int
}

type IRouter interface {
	//注册路由
	AddRouter(apiId int32, auth []int, handle HandleFunc)
	//路由处理
	Handle(cc ICConn, apiId int32, data []byte) rStatus.RInfo

	GetHandleMap() map[int]*Function
}
