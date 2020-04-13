package mManager

import (
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

const (
	MANAGER = "manager"
	SINGLEMODULENUM = "singleModuleNum"
	DEFAULTPORT = "DefaultPort"
)
