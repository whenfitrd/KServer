package minterface

type IManager interface {
	//启动
	Start()
	//关闭
	Stop()
	//添加模块路由
	AddRouter(moduleType int, apiId int32, handle HandleFunc, auth int)
	//注册模块
	Register(moduleType, ip, port string)
	//删除模块by port
	CancelPort(moduleType, port int)
	//删除一类模块
	CancelModule(moduleType int)
	//删除所有模块
	CancelAll()

	GetAllModule() map[int]map[int]IGameModule

	GetModuleByType(t int) map[int]IGameModule

	getUnusedPort(moduleType int) string
	//捕获ctrl+c来退出
	exitHandle()
}
