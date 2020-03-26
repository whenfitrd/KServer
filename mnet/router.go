package mnet

var HandleMap map[int]HandleFunc

type HandleFunc func([]byte)

type Router struct {
}

func (router *Router) Init() {
	HandleMap = make(map[int]HandleFunc)
}

func (router *Router) AddRouter(apiId int32, handle HandleFunc) {
	HandleMap[int(apiId)] = handle
}

func (router *Router) Handle(apiId int32, data []byte) {
	HandleMap[int(apiId)](data)
}
