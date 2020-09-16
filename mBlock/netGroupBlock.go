package mBlock

import (
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/mnet"
)

type NetGroupBlock struct {
	Block
	//网络组管理
	netGroupManager minterface.INetGroupManager
}

func (ngb *NetGroupBlock) GetNetGroupManager() minterface.INetGroupManager {
	return ngb.netGroupManager
}

func (ngb *NetGroupBlock) Config() {
	ngb.netGroupManager = mnet.GetNetGroupManager()
}
