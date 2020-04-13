package mManager

import (
	"fmt"
	"github.com/whenfitrd/KServer/mGame"
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/utils"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

type Manager struct {
	sync.Mutex
	//用于控制Manager的关闭
	managerChan chan bool
	//模块map 前int为模块类型后int为模块port
	ModuleMap map[int]map[int]minterface.IGameModule
	//注册端口号前int为模块类型后int为模块port(默认的port值延续点)
	NextPort map[int]int
	//白名单
	PortWhite map[int]chan int
	//起始端口号
	DefaultPort int
	//一种模块的数量
	SingleModuleNum int
	//模块的接口路由 模块类型  apiId
	Router map[int]map[int32]*minterface.Function
}

var m *Manager

func ApplyManager() *Manager {
	if m == nil {
		m = &Manager{
			managerChan: make(chan bool),
			Router: make(map[int]map[int32]*minterface.Function),
			NextPort: make(map[int]int),
			PortWhite: make(map[int]chan int),
			ModuleMap: make(map[int]map[int]minterface.IGameModule),
		}
	}
	m.Init()
	return m
}

func (m *Manager) Init() {
	//初始化log
	logger.Init()
	//加载配置文件
	ini.Load("config.ini")
	//从配置文件中获取并设置log文件
	logger.SetLogFile()

	//初始化defualtPort
	pstr, _ := ini.GetValue(MANAGER, DEFAULTPORT)
	p, e := strconv.Atoi(pstr)
	if e == nil {
		m.DefaultPort = p
	} else {
		m.DefaultPort = 50000
	}
	logger.Info("use default port", m.DefaultPort)

	//初始化singleModuleNum
	nstr, _ := ini.GetValue(MANAGER, SINGLEMODULENUM)
	n, e := strconv.Atoi(nstr)
	if e == nil {
		m.SingleModuleNum = n
	} else {
		m.SingleModuleNum = 10
	}
	logger.Info("SingleModuleNum: ", m.SingleModuleNum)
}

func (m *Manager) Start() {
	logger.Info("Start manager...")
	cmd := new(Cmd)
	rpc.Register(cmd)
	rpc.HandleHTTP()

	m.exitHandle()

	go func() {
		err := http.ListenAndServe(":50001", nil)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	select {
	case <-m.managerChan:
		break
	}
}

func (m *Manager) Stop() {
	logger.Info("Stop manager...")

	m.CancelAll()

	//关闭log
	if logger.Closed {
		return
	}
	logger.Close<- true
	<-logger.Clear

	m.managerChan <- true
}

func (m *Manager) AddRouter(moduleType int, apiId int32, handle minterface.HandleFunc, auth int) {
	_, ok := m.Router[moduleType]
	if !ok {
		m.Router[moduleType] = make(map[int32]*minterface.Function)
	}
	m.Router[moduleType][apiId] = &minterface.Function{
		Func: handle,
		Auth: auth,
	}
}

func (m *Manager) Register(moduleType int, ip string) {
	logger.Info("Register ...")
	port := m.getUnusedPort(moduleType)
	ms := mGame.ApplyModule(moduleType)

	_, ok := m.ModuleMap[moduleType]
	if !ok {
		m.ModuleMap[moduleType] = make(map[int]minterface.IGameModule)
	}
	port_i, e := strconv.Atoi(port)
	if e != nil {
		logger.Error("port error.", e.Error())
	}
	m.ModuleMap[moduleType][port_i] = ms

	for apiId, function := range m.Router[moduleType] {
		ms.GetServer().AddRouter(apiId, function)
	}

	logger.Info("module_%d_%s", moduleType, port)
	ms.Start(fmt.Sprintf("module_%d_%s", moduleType, port), ip, port)
}

func (m *Manager) CancelPort(moduleType, port int) {
	logger.Info("Cancel port...")
	_, ok := m.PortWhite[moduleType]
	if !ok {
		m.PortWhite[moduleType] = make(chan int, m.SingleModuleNum)
	}
	_, ok = m.ModuleMap[moduleType]
	if ok {
		_, o := m.ModuleMap[moduleType][port]
		if o {
			m.ModuleMap[moduleType][port].Stop()
			delete(m.ModuleMap[moduleType], port)
			m.PortWhite[moduleType] <- port
		}
	}
}

func (m *Manager) CancelModule(moduleType int) {
	logger.Info("Cancel module...")
	//清空对应的白名单
	delete(m.PortWhite, moduleType)
	//初始化对应的起始端口号
	m.NextPort[moduleType] = m.DefaultPort + moduleType * m.SingleModuleNum
	//关闭所有该种类的模块
	for _, s := range m.ModuleMap[moduleType] {
		s.Stop()
	}
	m.ModuleMap[moduleType] = make(map[int]minterface.IGameModule)
}

func (m *Manager) CancelAll() {
	logger.Info("Cancel all...")
	for mtype, _ := range m.ModuleMap {
		m.CancelModule(mtype)
	}
}

func (m *Manager) GetAllModule() map[int]map[int]minterface.IGameModule {
	return m.ModuleMap
}

func (m *Manager) GetModuleByType(t int) map[int]minterface.IGameModule {
	moduleM, ok := m.ModuleMap[t]
	if !ok {
		return nil
	}
	return moduleM
}

func (m *Manager) getUnusedPort(moduleType int) string {
	//从白名单中获取port
	for {
		_, ok := m.PortWhite[moduleType]
		if ok {
			if len(m.PortWhite[moduleType]) != 0 {
				port := <-m.PortWhite[moduleType]
				if utils.CheckPortIsUsed(port) {
					continue
				}
				return strconv.Itoa(port)
			}
			break
		}
		break
	}
	for {
		p, ok := m.NextPort[moduleType]
		if !ok {
			m.NextPort[moduleType] = m.DefaultPort + moduleType * m.SingleModuleNum
		} else {
			m.NextPort[moduleType] = p + 1
		}
		if utils.CheckPortIsUsed(m.NextPort[moduleType]) {
			continue
		}
		break
	}
	return strconv.Itoa(m.NextPort[moduleType])
}

var shutdownSignals = []os.Signal{os.Interrupt, os.Kill}

func (m *Manager) exitHandle() {
	//添加Ctrl+C的捕获处理
	logger.Info("ExitHandle")
	c := make(chan os.Signal, 1)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		logger.Info("Manager closing ...")
		go m.Stop()
		<-c
		<-c
		<-c
		logger.Warn("Force manager shutdown ...")
		os.Exit(1)
	}()
}