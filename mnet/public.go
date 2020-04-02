package mnet

import (
	"github.com/whenfitrd/KServer/global"
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/utils"
)

var g *global.GlobalObj
var logger *utils.Logger
var ngm minterface.INetGroupManager
var r minterface.IRouter

func init() {
	g = global.GetGObj()
	logger = utils.GetLogger()
	ngm = getNetGroupManager()
	r = getRouter()
}
