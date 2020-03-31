package mnet

//网络组管理
type NetGroupManager struct {
	NetGroups map[string]*NetGroup
}

//网络组
type NetGroup struct {
	Name string
	CCons map[string]*CConn
}

//添加链接至网络组
func (ngm *NetGroupManager)AddNetGroup(ccon *CConn, groupName string) {
	ng, ok := ngm.NetGroups[groupName]
	if !ok {
		ngm.NetGroups[groupName] = &NetGroup{
			Name: groupName,
			CCons: make(map[string]*CConn),
		}
		ngm.NetGroups[groupName].CCons[ccon.UID] = ccon
	} else {
		_, ok := ng.CCons[ccon.UID]
		if !ok {
			ng.CCons[ccon.UID] = ccon
		} else {
			ng.CCons[ccon.UID] = ccon
		}
	}
}

//删除网络组
func (ngm *NetGroupManager)DeleteNetGroup(groupName string) {
	_, ok := ngm.NetGroups[groupName]
	if ok {
		delete(ngm.NetGroups, groupName)
	}
}

//查找网络组
func (ngm *NetGroupManager)FindNetGroup(groupName string) *NetGroup {
	ng, ok := ngm.NetGroups[groupName]
	if ok {
		return ng
	} else {
		logger.Error("The group %s is not exist", groupName)
		return nil
	}
}

//离开网络组
func (ngm *NetGroupManager)LeaveNetGroup(ccon *CConn, groupName string) {
	ng, ok := ngm.NetGroups[groupName]
	if ok {
		delete(ng.CCons, ccon.UID)
	} else {
		logger.Error("The group %s is not exist", groupName)
	}
}

