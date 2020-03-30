package main

import (
	//"fmt"

	"github.com/whenfitrd/KServer/mnet"
)

func main() {
	s := mnet.ApplyServer("testServer", "localhost", "50000")
	s.Start()
}
