package mGame

import (
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/mnet"
)

type LoginModule struct {
	Server minterface.IServer
}

func (loginM *LoginModule)Start() {
	logger.Info("Start login module...")
	s := mnet.ApplyServer()
	loginM.Server = s
	s.SConfig("loginServer","0.0.0.0", "51000")
	s.Start()
}

func (loginM *LoginModule)Stop() {
	logger.Info("Stop login module...")
	loginM.Stop()
}
