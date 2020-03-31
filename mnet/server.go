package mnet

import (
	"github.com/whenfitrd/KServer/gObj"
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/utils"
	"net"
	"os"
	"os/signal"
)

type Server struct {
	Name string
	Ip   string
	Port string
}

var logger *utils.Logger
var gObject *gObj.GlobalObj

func ApplyServer(name, ip, port string) *Server {
	s := &Server{
		Name: name,
		Ip:   ip,
		Port: port,
	}

	s.Init()

	return s
}

func (s *Server) Init() {
	logger = utils.GetLogger()
	logger.Init()
	gObject = gObj.GetGObj()
	gObject.Init(GetNetGroupManager(), GetRouter())
}

func (s *Server) Start() {
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

func (s *Server) AddRouter(apiId int32, handle minterface.HandleFunc) {
	gObj.GetGObj().Router.GetHandleMap()[int(apiId)] = handle
}


