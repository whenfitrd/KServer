package mnet

type HandleFunc func(*CConn, []byte)

type Router struct {
	handleMap map[int]HandleFunc
}

func (router *Router) Init() {
	router.handleMap = make(map[int]HandleFunc)
}

func (router *Router) AddRouter(apiId int32, handle HandleFunc) {
	router.handleMap[int(apiId)] = handle
}

func (router *Router) Handle(cc *CConn, apiId int32, data []byte) {
	f, ok := router.handleMap[int(apiId)]
	if ok {
		f(cc, data)
	} else {
		logger.Error("Error apiId")
	}

}
