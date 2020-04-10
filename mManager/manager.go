package mManager

import (
	"fmt"
	"github.com/whenfitrd/KServer/mStruct"
	"net/http"
	"net/rpc"
)

type Manager struct {
	managerChan chan bool
}

func ApplyManager() *Manager {
	//默认值
	s := &Manager{
		managerChan: make(chan bool),
	}

	return s
}

func (m *Manager) Start() {
	logger.Info("Start manager...")
	cmd := new(mStruct.Cmd)
	rpc.Register(cmd)
	rpc.HandleHTTP()

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
	m.managerChan <- true
}