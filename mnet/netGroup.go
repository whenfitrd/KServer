package mnet

import (
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/rStatus"
)

var netGroupManager *NetGroupManager

func GetNetGroupManager() *NetGroupManager {
	if netGroupManager == nil {
		netGroupManager = &NetGroupManager{
			NetGroups: make(map[string]*NetGroup),
		}
	}
	return netGroupManager
}

//网络组管理
type NetGroupManager struct {
	NetGroups map[string]*NetGroup
}

//网络组
type NetGroup struct {
	Name string
	CCons map[string]minterface.ICConn
}

//创建网络组
func (ngm *NetGroupManager)CreateNetGroup(groupName string) rStatus.RInfo {
	ngm.NetGroups[groupName] = &NetGroup{
		Name: groupName,
		CCons: make(map[string]minterface.ICConn),
	}
	return rStatus.StatusOK
}

//添加链接至网络组
func (ngm *NetGroupManager)AddNetGroup(ccon minterface.ICConn, groupName string) rStatus.RInfo {
	ng, ok := ngm.NetGroups[groupName]
	if !ok {
		ngm.NetGroups[groupName] = &NetGroup{
			Name: groupName,
			CCons: make(map[string]minterface.ICConn),
		}
		ngm.NetGroups[groupName].CCons[ccon.GetUID()] = ccon
	} else {
		_, ok := ng.CCons[ccon.GetUID()]
		if !ok {
			ng.CCons[ccon.GetUID()] = ccon
		} else {
			//已经存在则覆盖
			ng.CCons[ccon.GetUID()] = ccon
		}
	}
	return rStatus.StatusOK
}

//删除网络组
func (ngm *NetGroupManager)DeleteNetGroup(groupName string) rStatus.RInfo {
	_, ok := ngm.NetGroups[groupName]
	if ok {
		delete(ngm.NetGroups, groupName)
		return rStatus.StatusOK
	} else {
		logger.Error("The group %s is not exist", groupName)
		return rStatus.StatusError
	}
}

//查找网络组
func (ngm *NetGroupManager)FindNetGroup(groupName string) (map[string]minterface.ICConn, rStatus.RInfo) {
	ng, ok := ngm.NetGroups[groupName]
	if ok {
		return ng.CCons, rStatus.StatusOK
	} else {
		logger.Error("The group %s is not exist", groupName)
		return nil, rStatus.StatusError
	}
}

//离开网络组
func (ngm *NetGroupManager)LeaveNetGroup(ccon minterface.ICConn, groupName string) rStatus.RInfo {
	ng, ok := ngm.NetGroups[groupName]
	if ok {
		delete(ng.CCons, ccon.GetUID())
		return rStatus.StatusOK
	} else {
		logger.Error("The group %s is not exist", groupName)
		return rStatus.StatusError
	}
}
