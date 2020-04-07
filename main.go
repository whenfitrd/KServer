package main

import (
	//"fmt"

	"github.com/whenfitrd/KServer/mnet"
)

func main() {
	s := mnet.ApplyServer()
	//s.SConfig("testServer", "0.0.0.0", "50000")
	s.Start()
	//s.LoadIni("config.ini")
}
