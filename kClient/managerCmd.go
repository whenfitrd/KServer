package main

import (
	"bufio"
	"fmt"
	"github.com/whenfitrd/KServer/mManager"
	"net/rpc"
	"strings"

	//"github.com/whenfitrd/KServer/mManager"
	//"net/rpc"
	"os"
	//"strings"
)

func main() {
	// 从标准输入流中接收输入数据
	//input := bufio.NewReader(os.Stdin)

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:50001")
	if err != nil {
		fmt.Println("dialing", err)
		return
	}

	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
		args := mManager.Args{
			Cmd:    os.Args[1],
			Params: os.Args[2:],
		}
		replay := mManager.Reply{}
		err := client.Call("Cmd.RpcReplay", args, &replay)
		if err != nil {
			fmt.Println("Cmd error:", err)
		}
		fmt.Println(replay.Message)
		return
	}

	//fmt.Printf("kServer# ")
	//// 逐行扫描
	////for input.Scan() {
	////str, err := input.ReadString(' ')
	////if err != nil {
	////	fmt.Printf("err" + str)
	////} else {
	////	fmt.Printf(str)
	////}
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {

		line := input.Text()

		// 输入exit时 结束
		if line == "exit" {
			break
		}

		line = strings.Trim(line, " ")

		cmdStr := strings.Split(line, " ")

		if len(cmdStr) == 0 {
		} else {
			args := mManager.Args{
				Cmd:    cmdStr[0],
				Params: cmdStr[1:],
			}
			replay := mManager.Reply{}
			err := client.Call("Cmd.RpcReplay", args, &replay)
			if err != nil {
				fmt.Println("Cmd error:", err)
			}
			fmt.Println(replay.Message)
		}

		fmt.Printf("kServer# ")
	}

}
