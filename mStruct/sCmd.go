package mStruct

import "fmt"

type Args struct {
	Cmd string
	Params []string
}

type Reply struct {
	Message string
}

type Cmd struct {
}

func (c *Cmd)RpcReplay(args *Args, reply *Reply) error {
	r := Reply{}
	if len(args.Params) == 1 {
		switch args.Cmd {
		case "start":
			switch args.Params[0] {
			case "loginModule":
				//启动loginModule
				r.Message = "Start loginModule ..."
			}
	case "stop":
			switch args.Params[0] {
			case "loginModule":
				//关闭loginModule
				r.Message = "Stop loginModule ..."
			}
		}
	} else {
		r.Message = fmt.Sprintf("Errror args about %s.", args.Cmd)
	}
	*reply = r
	return nil
}


