package mnet

import (
	"kserver/global"
	"kserver/utils"
	"net"
)

type CConn struct {
	TConn *net.TCPConn
	Router *Router
	BufChan chan []byte
}

func (cc *CConn) Init(tc *net.TCPConn, router *Router) {
	cc.TConn = tc
	cc.Router = router
	cc.BufChan = make(chan []byte, 16)
}

func (cc *CConn) Write() {}

func (cc *CConn) Read() {
	go cc.Handle()
	for {
		buf := make([]byte, 1024)
		_, err := cc.TConn.Read(buf)
		if err != nil {
			utils.GetLogger().Error(cc.TConn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		cc.BufChan<- buf
	}
}

func (cc *CConn) Handle() {
	//buffer := make([]byte, 1024)
	defer utils.HandlePanic()
	for {
		buf := <-cc.BufChan
		buffer := buf
		logger.Info("buffer len: ", len(buffer))
		msg := &Message{}
		msg.ParserHead(buffer[:global.MyMsgLen])
		buffer = buffer[global.MyMsgLen:]
		msg.MsgInfo = msg.ParserDataInfo(buffer[:global.MsgInfoLen]).(*MMsg)
		buffer = buffer[global.MsgInfoLen:]
		msg.Parser(buffer[:msg.MsgInfo.Length])
		logger.Info("buffer data: ", msg.MsgInfo.GetData())
		cc.Router.Handle(msg.MsgInfo.GetApiId(), msg.MsgInfo.GetData())
	}
}

func (cc *CConn) Close() {}
