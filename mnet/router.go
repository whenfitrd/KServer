package mnet

type HandleFunc func([]byte)

type Router struct {
	handleMap map[int]HandleFunc
}

func (router *Router) Init() {
	router.handleMap = make(map[int]HandleFunc)
}

func (router *Router) AddRouter(apiId int32, handle HandleFunc) {
	router.handleMap[int(apiId)] = handle
}

func (router *Router) Handle(apiId int32, data []byte) {
	router.handleMap[int(apiId)](data)
}
