package utils

import (
	"fmt"
	"github.com/whenfitrd/KServer/rStatus"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

type ILog interface {
	string() string
}

var logger *Logger

type Logger struct {
	sync.Mutex
	Name    string
	MsgChan chan *LogMsg
	Close   chan bool
	Clear   chan bool
	Closed  bool
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
			Closed: false,
		}
	}
	return logger
}

func (logger *Logger) Init() {
	go logger.RegisterCloseHandle()
	go logger.Start()
}

func (logger *Logger) SetLogFile() {
	if iniParser != nil {
		filePath, rst := iniParser.GetValue("base", "logFile")
		if rst != rStatus.StatusOK {
			filePath = "config.ini"
		}
		if filePath != "" {
			logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
			if err == nil {
				log.SetOutput(io.MultiWriter(os.Stderr,logFile))
			}
		}
	}
}

func (logger *Logger) RegisterCloseHandle() {
	//注册一个关闭处理的方法
	<-logger.Close
	logger.Warn("Logger start closing...")
	for {
		if len(logger.MsgChan) == 0 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	logger.Closed = true
	logger.Clear<- true
}

func (logger *Logger) PutMsg(msg *LogMsg) {
	//把msg放入缓冲池里
	if logger.Closed {
		return
	}
	logger.Lock()
	defer logger.Unlock()
	logger.MsgChan<- msg
}

func (logger *Logger) PopMsg() *LogMsg {
	return <-logger.MsgChan
}

//开启log
func (logger *Logger) Start() {
	for {
		msg := logger.PopMsg()
		levelString := ""
		switch msg.Level {
		case UNKNOWN:
			levelString = "[unknown]"
		case DEBUG:
			levelString = "[debug]"
		case INFO:
			levelString = "[info]"
		case WARN:
			levelString = "[warning]"
		case ERROR:
			levelString = "[error]"
		default:
			levelString = "[unknown]"
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

func (logger *Logger) debug(msg string, itf ...interface{}) {
	logger.Print(msg, DEBUG, itf)
}

func (logger *Logger) Info(msg string, itf ...interface{}) {
	logger.Print(msg, INFO, itf)
}

func (logger *Logger) Warn(msg string, itf ...interface{}) {
	logger.Print(msg, WARN, itf)
}

func (logger *Logger) Error(msg string, itf ...interface{}) {
	logger.Print(msg, ERROR, itf)
}

const (
	UNKNOWN = iota
	DEBUG
	INFO
	WARN
	ERROR
)
