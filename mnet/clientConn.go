package mnet

import (
	"github.com/whenfitrd/KServer/global"
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/utils"
	"net"
	"sync"
)

type CConn struct {
	sync.Mutex
	TConn *net.TCPConn
	BufChan chan []byte
	UID string
	Authorization int
}

func (cc *CConn) Init(tc *net.TCPConn) {
	cc.UID = utils.UniqueString()
	cc.TConn = tc
	cc.BufChan = make(chan []byte, 16)
	cc.Authorization = global.RVisitor
}

func (cc *CConn) Write(data []byte) {
	cc.Lock()
	defer cc.Unlock()
	_, err := cc.TConn.Write(data)
	if err != nil {
		logger.Error(cc.TConn.RemoteAddr().String(), " connection error: ", err)
		return
	}
}

func (cc *CConn) Read(router minterface.IRouter) {
	go cc.Handle(router)
	for {
		buf := make([]byte, 1024)
		_, err := cc.TConn.Read(buf)
		if err != nil {
			logger.Error(cc.TConn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		cc.BufChan<- buf
	}
}

func (cc *CConn) Handle(router minterface.IRouter) {
	defer utils.HandlePanic("clientConn handle close.")
	for {
		if buf, ok := <-cc.BufChan; ok {
			buffer := buf
			logger.Info("buffer len: ", len(buffer))
			m := utils.UnPackMsg(buffer, &Message{})
			router.Handle(cc, m.GetMsgInfo().GetApiId(), m.GetMsgInfo().GetData())
		} else {
			logger.Warn("Close Handle channel.")
			break
		}
	}
}

func (cc *CConn) Close() {
	cc.TConn.Close()
	for {
		if len(cc.BufChan) == 0 {
			close(cc.BufChan)
			break
		}
	}
}

func (cc *CConn) GetUID() string {
	return cc.UID
}

func (cc *CConn) UpdateAuth(auth int) {
	cc.Authorization = auth
}

func (cc *CConn) GetAuth() int {
	return cc.Authorization
}
