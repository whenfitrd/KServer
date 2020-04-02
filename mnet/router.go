package mnet

import "github.com/whenfitrd/KServer/minterface"

var router *Router

func getRouter() *Router {
	if router == nil {
		router = &Router{
			HandleMap: make(map[int]minterface.HandleFunc),
		}
	}
	return router
}

type Router struct {
	HandleMap map[int]minterface.HandleFunc
}

func (router *Router) AddRouter(apiId int32, handle minterface.HandleFunc) {
	router.HandleMap[int(apiId)] = handle
}

func (router *Router) Handle(cc minterface.ICConn, apiId int32, data []byte) {
	f, ok := router.HandleMap[int(apiId)]
	if ok {
		f(cc, data)
	} else {
		logger.Error("Error apiId")
	}
}

func (router *Router) GetHandleMap() map[int]minterface.HandleFunc {
	return router.HandleMap
}
