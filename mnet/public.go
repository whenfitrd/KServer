package mnet

import (
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/utils"
)

var ini *utils.IniParser
var logger *utils.Logger
var ngm minterface.INetGroupManager
var r minterface.IRouter
var pm minterface.IPlayerManager

func init() {
	logger = utils.GetLogger()
	ini = utils.GetIniParser()
	ngm = GetNetGroupManager()
	r = GetRouter()
	pm = GetPlayerManager()
}
