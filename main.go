package main

import (
	//"fmt"

	"kserver/mnet"
)

func main() {
	s := mnet.ApplyServer("testServer", "localhost", "50000")
	s.Start()
}
