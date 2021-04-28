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
	GroupMap map[string]int
	//客户端链接
	PConn minterface.ICConn
}

func (p *Player)Init(playerName string, cc minterface.ICConn) rStatus.RInfo {
	p.Name = playerName
	p.GroupMap = make(map[string]int)
	p.PConn = cc
	return rStatus.StatusOk
}

func (p *Player)CreateGroup(groupMap string) rStatus.RInfo {
	rst := ngm.CreateNetGroup(groupMap)
	if rst != rStatus.StatusOk {
		return rStatus.StatusError
	}
	p.GroupMap[groupMap] = global.Admin
	return rStatus.StatusOk
}

func (p *Player)AddToGroup(groupMap string) rStatus.RInfo {
	rst := ngm.AddNetGroup(p.PConn, groupMap)
	if rst != rStatus.StatusOk {
		return rStatus.StatusError
	}
	p.GroupMap[groupMap] = global.Member
	return rStatus.StatusOk
}

func (p *Player)LeaveNetGroup(groupMap string) rStatus.RInfo {
	rst := ngm.LeaveNetGroup(p.PConn, groupMap)
	if rst != rStatus.StatusOk {
		return rStatus.StatusError
	}
	delete(p.GroupMap, groupMap)
	return rStatus.StatusOk
}

func (p *Player)ListNetGroup() (rStatus.RInfo, map[string]int) {
	return rStatus.StatusOk, p.GroupMap
}

func (p *Player)KickOutPlayer(playerName, groupMap string) rStatus.RInfo {
	auth, ok := p.GroupMap[groupMap]
	if !ok {
		return rStatus.StatusError
	}
	if auth == global.Admin {
		player, rst := pm.FindPlayer(playerName)
		if rst != rStatus.StatusOk {
			return rStatus.StatusError
		}
		rst = player.LeaveNetGroup(groupMap)
		if rst != rStatus.StatusOk {
			return rStatus.StatusError
		}
		return rStatus.StatusOk
	} else {
		return rStatus.StatusOk
	}
}

func (p *Player)Exit() rStatus.RInfo {
	for gn, _ := range p.GroupMap {
		ngm.LeaveNetGroup(p.PConn, gn)
	}

	rst := pm.DeletePlayer(p.GetName())
	if rst != rStatus.StatusOk {
		return rStatus.StatusError
	}

	return rStatus.StatusOk
}

func (p *Player)GetName() string {
	return p.Name
}

func (p *Player)GetCConn() minterface.ICConn {
	return p.PConn
}
