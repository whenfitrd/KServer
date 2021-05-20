package minterface

type IMsg interface {
	//获取api优先级
	GetPriority() int32
	//获取APIID
	GetApiId() int32
	//获取数据长度
	GetLenth() int32
	//获取byte数据
	GetData() []byte
	//解析数据信息
	Parser([]byte)
}