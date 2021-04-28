package minterface

import "github.com/whenfitrd/KServer/rStatus"

type IPlayerManager interface {
	//创建玩家
	CreatePlayer(playerName string, cc ICConn) rStatus.RInfo
	//删除玩家
	DeletePlayer(playerName string) rStatus.RInfo
	//修改玩家信息
	UpdatePlayer(player IPlayer) rStatus.RInfo
	//查找玩家
	FindPlayer(playerName string) (IPlayer, rStatus.RInfo)
}

type IPlayer interface {
	//初始化玩家信息
	Init(playerName string, cc ICConn) rStatus.RInfo
	//创建网络组并加入
	CreateGroup(groupName string) rStatus.RInfo
	//加入网络组
	AddToGroup(groupName string) rStatus.RInfo
	//离开网络组
	LeaveNetGroup(groupName string) rStatus.RInfo
	//列出所属网络组名称和权限
	ListNetGroup() (rStatus.RInfo, map[string]int)
	//踢出网络组
	KickOutPlayer(playerName, groupName string) rStatus.RInfo
	//玩家退出
	Exit() rStatus.RInfo
	//获取玩家名
	GetName() string
	//获取客户端链接
	GetCConn() ICConn
}