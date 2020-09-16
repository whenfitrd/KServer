package mBlock

import (
	"github.com/whenfitrd/KServer/minterface"
)

type ServerBlock struct {
	Block
	Server minterface.IServer
}

func (sBlock *ServerBlock) Start(name, ip, port string) {
	logger.Info("Start server block...")
	sBlock.Server.SConfig(name,ip, port)
	sBlock.Server.Start()
}

func (sBlock *ServerBlock) Stop() {
	logger.Info("Stop server block...")
	sBlock.Server.Stop()
}

func (sBlock *ServerBlock) GetServer() minterface.IServer {
	return sBlock.Server
}
