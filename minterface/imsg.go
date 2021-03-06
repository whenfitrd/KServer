package minterface

import "github.com/whenfitrd/KServer/rStatus"

type IMsg interface {
	GetType() int8

	GetMsgParser(msgType int8) IMMsg

	GetMsgInfo() IMMsg

	ParserHead(data []byte) rStatus.RStatus

	ParserDataInfo(data []byte)

	Parser(data []byte)
}

type IMMsg interface {
	//获取数据长度
	GetLength() int32
	//获取APIID
	GetApiId() int32
	//获取byte数据
	GetData() []byte
	//接续数据
	Parser([]byte)
	//解析数据信息(长度和apiId)
	ParserDataInfo([]byte)
}
