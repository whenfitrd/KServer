package main

import (
	"encoding/json"
	"fmt"
	"github.com/whenfitrd/KServer/global"
	"github.com/whenfitrd/KServer/mManager"
	"github.com/whenfitrd/KServer/minterface"
)

type Test struct{
	Info string
	L int
}

func main() {
	mgr := mManager.ApplyManager()
	mgr.AddRouter(global.ServerBlock, 1, test, global.RAll)
	mgr.Start()
}

func test(cc minterface.ICConn, data []byte) {
	var s Test

	err:=json.Unmarshal([]byte(data), &s)
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println("=================test api======", s.Info, s.L)
}
