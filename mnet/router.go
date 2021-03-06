package mnet

import (
	"github.com/whenfitrd/KServer/global"
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/rStatus"
)

var router *Router

func GetRouter() *Router {
	if router == nil {
		router = &Router{
			HandleMap: make(map[int]*minterface.Function),
			Auth: global.RAll,
		}
	}
	return router
}

type Router struct {
	HandleMap map[int]*minterface.Function
	Auth int
}

func (router *Router) SetAuth(auth int) {
	router.Auth = auth
}

func (router *Router) AddRouter(apiId int32, handle minterface.HandleFunc) {
	router.HandleMap[int(apiId)] = &minterface.Function{
		Func: handle,
		Auth: router.Auth,
	}
}

func (router *Router) Handle(cc minterface.ICConn, apiId int32, data []byte) rStatus.RInfo {
	f, ok := router.HandleMap[int(apiId)]
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

func (router *Router) GetHandleMap() map[int]*minterface.Function {
	return router.HandleMap
}
