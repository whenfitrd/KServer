package mnet

import (
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/rStatus"
)

type Router struct {
	HandleMap map[int]*minterface.Function
	Auth int
}

//func (router *Router) SetAuth(auth int) {
//	router.Auth = auth
//}
//
//func (router *Router) AddRouter(apiId int32, handle minterface.HandleFunc) {
//	router.HandleMap[int(apiId)] = &minterface.Function{
//		Func: handle,
//		Auth: router.Auth,
//	}
//}

func (rt *Router) AddRouter(apiId int32, handle *minterface.Function) {
	rt.HandleMap[int(apiId)] = handle
}

func (rt *Router) Handle(cc minterface.ICConn, apiId int32, data []byte) rStatus.RInfo {
	f, ok := rt.HandleMap[int(apiId)]
	if !ok {
		logger.Error("Error apiId")
		return rStatus.StatusError
	}
	//简单的接口权限检测
	if CheckAuth(cc.GetAuth(), f.Auth) {
		f.Func(cc, data)
		return rStatus.StatusOk
	}
	return rStatus.ApiAuthError
}

func (rt *Router) GetHandleMap() map[int]*minterface.Function {
	return rt.HandleMap
}
