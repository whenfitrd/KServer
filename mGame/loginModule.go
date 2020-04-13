package mGame

import (
	"github.com/whenfitrd/KServer/minterface"
)

type LoginModule struct {
	Server minterface.IServer
}

func (loginM *LoginModule) Start(name, ip, port string) {
	logger.Info("Start login module...")
	loginM.Server.SConfig(name,ip, port)
	loginM.Server.Start()
}

func (loginM *LoginModule) Stop() {
	logger.Info("Stop login module...")
	loginM.Server.Stop()
}

func (loginM *LoginModule) GetServer() minterface.IServer {
	return loginM.Server
}
