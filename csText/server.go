package main

import (
	//"encoding/json"
	"fmt"
	"github.com/whenfitrd/KServer/global"
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/mnet"
	"github.com/whenfitrd/KServer/pb"

	"github.com/golang/protobuf/proto"
)

func main() {
	s := mnet.ApplyServer()
	s.Init()
	s.SetAuth(global.RVisitor)
	s.AddRouter(1, test)
	s.Start()
}

func test(cc minterface.ICConn, data []byte) {
	var s pb.Base

	fmt.Println(data)
	proto.UnmarshalMerge(data, &s)

	fmt.Println("=================test api======", s.Info, s.Head)

	//err:=json.Unmarshal([]byte(data), &s)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//
	//fmt.Println("=================test api======", s.Info, s.L)
}
