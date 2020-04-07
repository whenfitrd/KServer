package mnet

import (
	"github.com/whenfitrd/KServer/global"
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/rStatus"
)

type Player struct {
	//名称
	Name string
	//所属网络组名称和权限
	GroupNames map[string]int
	//客户端链接
	PConn minterface.ICConn
}

func (p *Player)Init(playerName string, cc minterface.ICConn) rStatus.RInfo {
	p.Name = playerName
	p.GroupNames = make(map[string]int)
	p.PConn = cc
	return rStatus.StatusOK
}

func (p *Player)CreateGroup(groupName string) rStatus.RInfo {
	rst := ngm.CreateNetGroup(groupName)
	if rst != rStatus.StatusOK {
		return rStatus.StatusError
	}
	p.GroupNames[groupName] = global.Admin
	return rStatus.StatusOK
}

func (p *Player)AddToGroup(groupName string) rStatus.RInfo {
	rst := ngm.AddNetGroup(p.PConn, groupName)
	if rst != rStatus.StatusOK {
		return rStatus.StatusError
	}
	p.GroupNames[groupName] = global.Member
	return rStatus.StatusOK
}

func (p *Player)LeaveNetGroup(groupName string) rStatus.RInfo {
	rst := ngm.LeaveNetGroup(p.PConn, groupName)
	if rst != rStatus.StatusOK {
		return rStatus.StatusError
	}
	delete(p.GroupNames, groupName)
	return rStatus.StatusOK
}

func (p *Player)KickOutPlayer(playerName, groupName string) rStatus.RInfo {
	auth, ok := p.GroupNames[groupName]
	if !ok {
		return rStatus.StatusError
	}
	if auth == global.Admin {
		player, rst := pm.FindPlayer(playerName)
		if rst != rStatus.StatusOK {
			return rStatus.StatusError
		}
		rst = player.LeaveNetGroup(groupName)
		if rst != rStatus.StatusOK {
			return rStatus.StatusError
		}
		return rStatus.StatusOK
	} else {
		return rStatus.StatusOK
	}
}

func (p *Player)Exit() rStatus.RInfo {
	for gn, _ := range p.GroupNames {
		ngm.LeaveNetGroup(p.PConn, gn)
	}

	rst := pm.DeletePlayer(p.GetName())
	if rst != rStatus.StatusOK {
		return rStatus.StatusError
	}

	return rStatus.StatusOK
}

func (p *Player)GetName() string {
	return p.Name
}

func (p *Player)GetCConn() minterface.ICConn {
	return p.PConn
}
