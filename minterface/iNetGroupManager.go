package minterface

import "github.com/whenfitrd/KServer/rStatus"

type INetGroupManager interface {
	//添加网络组
	AddNetGroup(ccon ICConn, groupName string) rStatus.RInfo
	//删除网络组
	DeleteNetGroup(groupName string) rStatus.RInfo
	//查找网络组
	FindNetGroup(groupName string) (map[string]ICConn, rStatus.RInfo)
	//离开网络组
	LeaveNetGroup(ccon ICConn, groupName string) rStatus.RInfo
}
