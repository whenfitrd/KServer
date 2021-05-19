package mnet

import (
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/utils"
)

var ini *utils.CParser
var logger *utils.Logger
var ngm minterface.INetGroupManager
var r minterface.IRouter
var pm minterface.IPlayerManager

func init() {
	logger = utils.GetLogger()
	ini = utils.GetConfigParser()
	ngm = GetNetGroupManager()
	r = GetRouter()
	pm = GetPlayerManager()
}
