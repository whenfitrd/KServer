package kNet

import (
	"github.com/whenfitrd/KServer/utils"
)

var ini *utils.IniParser
var logger *utils.Logger

func init() {
	logger = utils.GetLogger()
	ini = utils.GetIniParser()
}
