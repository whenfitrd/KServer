package main

import (
	//"fmt"

	"encoding/json"
	"fmt"
	"kserver/mnet"
)

type Test struct{
	Info string
	L int
}

//type SliceMock struct {
//	addr uintptr
//	len int
//	cap int
//}

func main() {
	s := mnet.ApplyServer("testServer", "localhost", "50000")
	s.Router.AddRouter(1, test)
	s.Start()
}

func test(data []byte) {
	fmt.Printf("=================test api================\n")
	//t := &Test{}
	//t := Test{}
	//buf := &bytes.Buffer{}
	//err := binary.Read(buf, binary.BigEndian, &t)
	//if err != nil {
	//	panic(err)
	//}

	//var t *Test = *(**Test)(unsafe.Pointer(&data))
	//a := (*Test)(unsafe.Pointer(&data))
	//Len := unsafe.Sizeof(&data)
	//testBytes := &SliceMock{
	//	addr: uintptr(unsafe.Pointer(&data)),
	//	cap: int(Len),
	//	len: int(Len),
	//}
	//d := *(*[]byte)(unsafe.Pointer(testBytes))
	var s Test
	err:=json.Unmarshal([]byte(data), &s)
	if err!=nil{
		fmt.Println(err)
	}

	//var a *Test = *(**Test)(unsafe.Pointer(&data))
	//var tt Test = *a


	fmt.Println("=================test api======", s.Info, s.L)
}
