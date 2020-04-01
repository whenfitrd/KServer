package main

import (
	"encoding/json"
	"fmt"
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/mnet"
)

type Test struct{
	Info string
	L int
}

func main() {
	s := mnet.ApplyServer("testServer", "localhost", "50000")
	s.AddRouter(1, test)
	s.Start()
}

func test(cc minterface.ICConn, data []byte) {
	var s Test

	err:=json.Unmarshal([]byte(data), &s)
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println("=================test api======", s.Info, s.L)
}
