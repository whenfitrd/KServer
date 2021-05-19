package utils

import (
	"bufio"
	"github.com/whenfitrd/KServer/rStatus"
	"log"
	"os"
	"strings"
	"sync"
)

type CParser struct {
	FileName string
	IniMap map[string]map[string]string
}

var cParser *CParser
var onceParser sync.Once

func GetConfigParser() *CParser {
	onceParser.Do(func() {
		cParser = &CParser{
			FileName: "",
			IniMap: make(map[string]map[string]string),
		}
	})
	return cParser
}

//加载配置文件
func (config *CParser)Load(fileName string) rStatus.RInfo {
	config.FileName = fileName
	config.IniMap = make(map[string]map[string]string)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return rStatus.StatusError
	}
	defer file.Close()

	tag := "defalut"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		//去除字符串前后的空格
		line := strings.Trim(l, " ")

		//处理空行，使用//#;作为注释关键字
		idx := strings.Index(line, ";")
		if idx >= 0 {
			line = line[:idx]
		}
		idx = strings.Index(line, "#")
		if idx >= 0 {
			line = line[:idx]
		}
		idx = strings.Index(line, "//")
		if idx >= 0 {
			line = line[:idx]
		}
		if len(line) == 0 {
			continue
		}

		lidx := strings.Index(line, "[")
		ridx := strings.Index(line, "]")
		if lidx > -1 && ridx > -1 {
			line = line[lidx+1:ridx]
			tag = line
			config.IniMap[tag] = make(map[string]string)
		} else if lidx < 0 && ridx < 0  {
			cidx := strings.Index(line, "=")
			if cidx == -1 {
				config.IniMap[tag][strings.Trim(line, " ")] = ""
			} else {
				config.IniMap[tag][strings.Trim(line[:cidx], " ")] = strings.Trim(line[cidx+1:], " ")
			}
		}
	}
	return rStatus.StatusOk
}


func (config *CParser)GetValue(tag, name string) (string, rStatus.RInfo) {
	m, ok := config.IniMap[tag]
	if !ok {
		return "", rStatus.StatusError
	}
	v, ok := m[name]
	if !ok {
		return "", rStatus.StatusError
	}
	return v, rStatus.StatusOk
}