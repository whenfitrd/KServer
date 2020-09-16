package kNet

import (
	"net"
)

type NetConnect struct {
	Name string
	Ip   string
	Port string
	serverChan chan bool
	listenerCloseFlag bool
}

func (nc *NetConnect) Config(name, ip, port string) {
	nc.Name = name
	nc.Ip = ip
	nc.Port = port
}

func (nc *NetConnect) Start() {
	logger.Info("Starting the netConnect ...")

	go nc.AcceptConnect()
}

func (nc *NetConnect) Stop() {
	logger.Info("Stop netConnect...")
	nc.listenerCloseFlag = true
	nc.serverChan <- true
}

func (nc *NetConnect) close(listener *net.TCPListener) {
	<-nc.serverChan
	logger.Warn("Stop listener ...")

	if err := listener.Close(); err != nil {
		logger.Error(err.Error())
	}
}

func (nc *NetConnect) AcceptConnect() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", nc.Ip+":"+nc.Port)
	if err != nil {
		logger.Error("Error ip, err: " + err.Error())
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		logger.Error("Error listening, err: " + err.Error())
		return
	}

	go nc.close(listener)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			if nc.listenerCloseFlag {
				logger.Info("Close lisener.")
			} else {
				logger.Error("Error accepting, err: " + err.Error())
			}
			break
		}
		go nc.ConnectHandle(conn)
	}

	logger.Info("Connect had been closed.")
}
