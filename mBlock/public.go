package mBlock

import (
	"github.com/whenfitrd/KServer/global"
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/mnet"
	"github.com/whenfitrd/KServer/utils"
)

var ini *utils.IniParser
var logger *utils.Logger
var ngm minterface.INetGroupManager
var pm minterface.IPlayerManager

func init() {
	logger = utils.GetLogger()
	ini = utils.GetIniParser()
	ngm = mnet.GetNetGroupManager()
	pm = mnet.GetPlayerManager()
}


func ApplyBlock(moduleType int) minterface.IBlock {
	switch moduleType {
	case global.ServerBlock:
		return &ServerBlock{
			Server: mnet.ApplyServer(),
		}
	}
	return nil
}
