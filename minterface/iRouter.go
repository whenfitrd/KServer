package minterface

type HandleFunc func(ICConn, []byte)

type IRouter interface {
	//注册路由
	AddRouter(apiId int32, handle HandleFunc)
	//路由处理
	Handle(cc ICConn, apiId int32, data []byte)

	GetHandleMap() map[int]HandleFunc
}
