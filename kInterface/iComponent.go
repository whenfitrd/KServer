package kInterface

//组件
type IComponent interface {
	//激活
	Active()
	//冻结
	Freeze()
	//获取块对象
	GetComponent() IComponent
}

//交互组件-与前段交互并转发操作给后续组件处理
type ICInteractive interface {
	IComponent
	//启动
	Start()
	//停止
	Stop()
}
