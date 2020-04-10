package main

import (
	"github.com/whenfitrd/KServer/mManager"
)

func main() {
	//s := mnet.ApplyServer()
	////s.SConfig("testServer", "0.0.0.0", "50000")
	//s.Start()
	//s.LoadIni("config.ini")

	mgr := mManager.ApplyManager()
	mgr.Start()
}
