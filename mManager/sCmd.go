package mManager

import (
	"fmt"
	"github.com/whenfitrd/KServer/global"
	"strconv"
)

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
	switch args.Cmd {
	case "start":
		if len(args.Params) == 1 {
			switch args.Params[0] {
			case "loginModule":
				logger.Info("start loginModule")
				go m.Register(global.ServerBlock, "0.0.0.0")
				r.Message = fmt.Sprintf("Start a loginModule ...")
			}
		}
	case "stop":
		if len(args.Params) == 1 {
			switch args.Params[0] {
			case "all":
				go m.CancelAll()
				r.Message = fmt.Sprintf("Stop all modules...")
			case "loginModule":
				go m.CancelModule(global.ServerBlock)
				r.Message = fmt.Sprintf("Stop loginModules...")
			}
		} else if len(args.Params) >= 2 {
			switch args.Params[0] {
			case "loginModule":
				port, e := strconv.Atoi(args.Params[1])
				if e != nil {
					r.Message = "Error params after module."
				}
				go m.CancelPort(global.ServerBlock, port)
				r.Message = fmt.Sprintf("Stop loginModules %d", port)
			}
		}
	case "get":
		if len(args.Params) == 1 {
			switch args.Params[0] {
			case "modules":
				msgstr := "moduleType port"
				for mtype, v := range m.ModuleMap {
					for port, _ := range v {
						msgstr += fmt.Sprintf("\n%s %d", moduleMap[mtype], port)
					}
				}
				r.Message = msgstr
			}
		}
	default:
		r.Message = fmt.Sprintf("Errror args about %s.", args.Cmd)
	}
	*reply = r
	return nil
}

var moduleMap = map[int]string{
	global.ServerBlock: "loginModule",
}

