package utils

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

type ILog interface {
	string() string
}

var logger *Logger

type Logger struct {
	Name    string
	MsgChan chan *LogMsg
	Close   chan bool
	Clear   chan bool
}

type LogMsg struct {
	Level int
	FPath string
	Line  int
	Msg   string
	Itf   []interface{}
}

func GetLogger() *Logger {
	if logger == nil {
		logger = &Logger{
			Name:    "logger",
			MsgChan: make(chan *LogMsg, 256),
			Close: make(chan bool),
			Clear: make(chan bool),
		}
	}
	return logger
}

func (logger *Logger) Init() {
	go logger.PrintLog()
	go logger.RegisterCloseHandle()
}

func (logger *Logger) RegisterCloseHandle() {
	//var b bool
	<-logger.Close
	logger.Warn("Logger start closing...")
	//if b {
	for {
		if len(logger.MsgChan) == 0 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	logger.Clear<- true
	//}
}

func (logger *Logger) PutMsg(msg *LogMsg) {
	logger.MsgChan<- msg
}

func (logger *Logger) PopMsg() *LogMsg {
	return <-logger.MsgChan
}

func (logger *Logger) PrintLog() {
	for {
		msg := logger.PopMsg()
		levelString := ""
		switch msg.Level {
		case 1:
			levelString = "[info]"
		case 2:
			levelString = "[warning]"
		case 3:
			levelString = "[error]"
		default:
			levelString = "[unknow]"
		}
		log.SetPrefix(levelString)
		loginfo := fmt.Sprintf("%s:%d %s", msg.FPath, msg.Line, msg.Msg)
		if len(msg.Itf) != 0 {
			extraInfo := fmt.Sprint("", msg.Itf)
			loginfo = loginfo + " " + extraInfo[1:len(extraInfo)-1]
		}
		log.Println(loginfo)
	}
}

func (logger *Logger) Print(m string, level int, itf []interface{}) {
	//for _, i := range itfs {
	//	itf = append(itf, i)
	//}
	_, file, line, ok := runtime.Caller(2)
	if ok {
		msg := &LogMsg{
			Level: level,
			FPath: file,
			Line:  line,
			Msg:   m,
			Itf:   itf,
		}
		logger.PutMsg(msg)
	}
}

func (logger *Logger) Info(msg string, itf ...interface{}) {
	logger.Print(msg, 1, itf)
}

func (logger *Logger) Warn(msg string, itf ...interface{}) {
	logger.Print(msg, 2, itf)
}

func (logger *Logger) Error(msg string, itf ...interface{}) {
	logger.Print(msg, 3, itf)
}
