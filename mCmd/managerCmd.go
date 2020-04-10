package main

import (
	"bufio"
	"fmt"
	"github.com/whenfitrd/KServer/mStruct"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	// 从标准输入流中接收输入数据
	input := bufio.NewScanner(os.Stdin)

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:50001")
	if err != nil {
		fmt.Println("dialing", err)
	}

	fmt.Printf("kServer# ")
	// 逐行扫描
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
			args := mStruct.Args{
				Cmd:    cmdStr[0],
				Params: cmdStr[1:],
			}
			replay := mStruct.Reply{}
			err := client.Call("Cmd.RpcReplay", args, &replay)
			if err != nil {
				fmt.Println("Cmd error:", err)
			}
			fmt.Println(replay.Message)
		}

		fmt.Printf("kServer# ")
	}

}
