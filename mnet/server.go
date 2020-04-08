package mnet

import (
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/rStatus"
	"net"
	"os"
	"os/signal"
)

type Server struct {
	Name string
	Ip   string
	Port string
}

func ApplyServer() *Server {
	//默认值
	s := &Server{
		Name: "testServer",
		Ip:   "localhost",
		Port: "50000",
	}

	return s
}

func (s *Server) SConfig(name, ip, port string) {
	s.Name = name
	s.Ip = ip
	s.Port = port
}

func (s *Server) LoadIni(fileName string) rStatus.RInfo {
	sts := ini.Load(fileName)
	if sts != rStatus.StatusOk {
		return rStatus.StatusError
	}
	logger.SetLogFile()
	return rStatus.StatusOk
}

func (s *Server) Init() {
	logger.Init()
	s.LoadIni("config.ini")
}

func (s *Server) Start() {
	s.Init()
	logger.Info("Starting the server ...")

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

	//Panic2Error()

	s.ExitHandle()

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			logger.Error("Error accepting, err: " + err.Error())
			return
		}
		go s.ConnectHandle(conn)
	}
}

func (s *Server) Stop() {
	s.LoggerClose()
	os.Exit(1)
}

func (s *Server) LoggerClose() {
	//关闭log
	if logger.Closed {
		return
	}
	logger.Close<- true
	<-logger.Clear
}

func (s *Server) ConnectHandle(conn *net.TCPConn) (err error) {
	//defer utils.HandlePanic()

	//添加路由处理
	cc := &CConn{}
	cc.Init(conn)
	cc.Read()

	//msg := &Message{}
	//msg.Parser(conn)
	//s.Router.Handle(msg.MsgInfo.GetApiId(), msg.MsgInfo.GetData())
	return nil
}

var shutdownSignals = []os.Signal{os.Interrupt, os.Kill}

func (s *Server) ExitHandle() {
	//添加Ctrl+C的捕获处理
	logger.Info("ExitHandle")
	c := make(chan os.Signal, 1)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		logger.Info("Server closing ...")
		go s.Stop()
		<-c
		<-c
		<-c
		logger.Warn("Force server shutdown ...")
		os.Exit(1)
	}()
}

func Panic2Error() (err error) {
	//panic(-1)
	return nil
}

func (s *Server) AddRouter(apiId int32, auth []int, handle minterface.HandleFunc)  {
	r.AddRouter(apiId, auth, handle)
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


