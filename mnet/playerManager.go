package mnet

import (
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/rStatus"
	"sync"
)

var playerManager *PlayerManager

func GetPlayerManager() *PlayerManager {
	if playerManager == nil {
		playerManager = &PlayerManager{
			PlayerList: make(map[string]minterface.IPlayer),
		}
	}
	return playerManager
}

type PlayerManager struct {
	sync.Mutex
	PlayerList map[string]minterface.IPlayer
}

func (pm *PlayerManager)CreatePlayer(playerName string, cc minterface.ICConn) rStatus.RInfo {
	pm.Lock()
	defer pm.Unlock()
	p := &Player{}
	p.Init(playerName, cc)
	pm.PlayerList[p.GetName()] = p
	return rStatus.StatusOk
}

func (pm *PlayerManager)DeletePlayer(playerName string) rStatus.RInfo {
	delete(pm.PlayerList, playerName)
	return rStatus.StatusOk
}

func (pm *PlayerManager)UpdatePlayer(player minterface.IPlayer) rStatus.RInfo {
	pm.Lock()
	defer pm.Unlock()
	_, rst := pm.FindPlayer(player.GetName())
	if rst != rStatus.StatusOk {
		return rStatus.StatusError
	}
	pm.PlayerList[player.GetName()] = player
	return rStatus.StatusError
}

func (pm *PlayerManager)FindPlayer(playerName string) (minterface.IPlayer, rStatus.RInfo) {
	p, rst := pm.PlayerList[playerName]
	if rst {
		return p, rStatus.StatusOk
	}
	return nil, rStatus.StatusError
}
