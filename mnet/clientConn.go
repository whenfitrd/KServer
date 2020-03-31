package mnet

import (
	"github.com/whenfitrd/KServer/global"
	"github.com/whenfitrd/KServer/utils"
	"net"
	"sync"
)

type CConn struct {
	sync.Mutex
	TConn *net.TCPConn
	Router *Router
	BufChan chan []byte
	UID string
}

func (cc *CConn) Init(tc *net.TCPConn, router *Router) {
	cc.UID = utils.UniqueString()
	cc.TConn = tc
	cc.Router = router
	cc.BufChan = make(chan []byte, 16)
	cc.Read()
}

func (cc *CConn) Write(data []byte) {
	_, err := cc.TConn.Write(data)
	if err != nil {
		utils.GetLogger().Error(cc.TConn.RemoteAddr().String(), " connection error: ", err)
		return
	}
}

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
	defer utils.HandlePanic("clientConn")
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
		cc.Router.Handle(cc, msg.MsgInfo.GetApiId(), msg.MsgInfo.GetData())
	}
}

func (cc *CConn) Close() {}
