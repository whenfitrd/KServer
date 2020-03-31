package gObj

import (
	"github.com/whenfitrd/KServer/minterface"
)

var gObject *GlobalObj

func GetGObj() *GlobalObj {
	if gObject == nil {
		gObject = &GlobalObj{}
	}
	return gObject
}

type GlobalObj struct {
	NetGroupManager minterface.INetGroupManager
	Router minterface.IRouter
}

func (gObject *GlobalObj)Init(ngm minterface.INetGroupManager, router minterface.IRouter) {
	gObject.NetGroupManager = ngm
	gObject.Router = router
}
