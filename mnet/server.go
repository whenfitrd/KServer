package mnet

import (
	"github.com/whenfitrd/KServer/global"
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/rStatus"
	"net"
	"sync"
)

type Server struct {
	Name string
	Ip   string
	Port string
	Router minterface.IRouter
	serverChan chan bool
	wg sync.WaitGroup
}

func ApplyServer() *Server {
	//默认值
	s := &Server{
		Name: "testServer",
		Ip:   "localhost",
		Port: "50000",
		Router: &Router{},
		serverChan: make(chan bool),
		wg: sync.WaitGroup{},
	}

	s.Init()
	return s
}

func (s *Server) SConfig(name, ip, port string) {
	s.Name = name
	s.Ip = ip
	s.Port = port
}

func (s *Server) Init() {
	//初始化路由
	s.Router = &Router{
		HandleMap: make(map[int]*minterface.Function),
		Auth: global.RAll,
	}
}

func (s *Server) Start() {
	logger.Info("Starting the server ...")

	go s.AcceptConnect()
}

func (s *Server) Stop() {
	logger.Info("Stop server...")
	s.serverChan <- true
}

func (s *Server) AcceptConnect() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", s.Ip+":"+s.Port)
	if err != nil {
		logger.Error("Error ip, err: " + err.Error())
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		logger.Error("Error listening, err: " + err.Error())
		return
	}

	go func() {
		<-s.serverChan
		logger.Warn("Stop listener ...")
		if err = listener.Close(); err != nil {
			logger.Error(err.Error())
		}
	}()

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			logger.Error("Error accepting, err: " + err.Error())
			return
		}
		s.wg.Add(1)
		go s.ConnectHandle(conn)
	}
}

func (s *Server) ConnectHandle(conn *net.TCPConn) (err error) {
	defer s.wg.Done()
	defer conn.Close()
	//defer utils.HandlePanic()

	//添加路由处理
	cc := &CConn{}
	cc.Init(conn)
	cc.Read(s.Router)

	//msg := &Message{}
	//msg.Parser(conn)
	//s.Router.Handle(msg.MsgInfo.GetApiId(), msg.MsgInfo.GetData())
	return nil
}

func Panic2Error() (err error) {
	//panic(-1)
	return nil
}

//func (s *Server) SetAuth(auth int) {
//	r.SetAuth(auth)
//}

//func (s *Server) AddRouter(apiId int32, handle minterface.HandleFunc)  {
//	r.AddRouter(apiId, handle)
//}
func (s *Server) AddRouter(apiId int32, handle *minterface.Function) {
	s.Router.AddRouter(apiId, handle)
}

func (s *Server) WriteToGroup(data []byte, groupName string) rStatus.RInfo {
	group, sts := ngm.FindNetGroup(groupName)
	if sts == rStatus.StatusOk {
		for _, conns := range group {
			conns.Write(data)
		}
		return rStatus.StatusOk
	} else {
		return rStatus.StatusError
	}
}


