package mnet

import (
	"github.com/whenfitrd/KServer/gObj"
	"github.com/whenfitrd/KServer/msg"
	"github.com/whenfitrd/KServer/utils"
	"net"
	"sync"
)

type CConn struct {
	sync.Mutex
	TConn *net.TCPConn
	BufChan chan []byte
	UID string
}

func (cc *CConn) Init(tc *net.TCPConn) {
	cc.UID = utils.UniqueString()
	cc.TConn = tc
	cc.BufChan = make(chan []byte, 16)
}

func (cc *CConn) Write(data []byte) {
	cc.Lock()
	defer cc.Unlock()
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
	defer utils.HandlePanic("clientConn")
	for {
		buf := <-cc.BufChan
		buffer := buf
		logger.Info("buffer len: ", len(buffer))
		m := msg.UnPackMsg(buffer)
		gObj.GetGObj().Router.Handle(cc, m.MsgInfo.GetApiId(), m.MsgInfo.GetData())
	}
}

func (cc *CConn) Close() {
	cc.TConn.Close()
}

func (cc *CConn) GetUID() string {
	return cc.UID
}
