package mnet

import (
	"kserver/utils"
	"net"
	"os"
	"os/signal"
)

type Server struct {
	Name string
	Ip   string
	Port string
	Router *Router
}

var logger *utils.Logger

func ApplyServer(name, ip, port string) *Server {
	s := &Server{
		Name: name,
		Ip:   ip,
		Port: port,
		Router: &Router{},
	}

	s.Init()

	return s
}

func (s *Server) Init() {
	s.Router.Init()
}

func (s *Server) Start() {
	logger = utils.GetLogger()
	logger.Init()
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
	logger.Close<- true
	<-logger.Clear
}

func (s *Server) ConnectHandle(conn *net.TCPConn) (err error) {
	defer utils.HandlePanic()

	cc := &CConn{}
	cc.Init(conn, s.Router)
	cc.Read()

	//msg := &Message{}
	//msg.Parser(conn)
	//s.Router.Handle(msg.MsgInfo.GetApiId(), msg.MsgInfo.GetData())
	return nil
}

var shutdownSignals = []os.Signal{os.Interrupt, os.Kill}

func (s *Server) ExitHandle() {
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


