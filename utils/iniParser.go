package utils

import (
	"bufio"
	"github.com/whenfitrd/KServer/rStatus"
	"log"
	"os"
	"strings"
)

var iniParser *IniParser

func GetIniParser() *IniParser {
	if iniParser == nil {
		iniParser = &IniParser{
			FileName: "",
			IniMap: make(map[string]map[string]string),
		}
	}
	return iniParser
}

type IniParser struct {
	FileName string
	IniMap map[string]map[string]string
}

func (ini *IniParser)Load(fileName string) rStatus.RInfo {
	ini.FileName = fileName
	ini.IniMap = make(map[string]map[string]string)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return rStatus.StatusError
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//只取分号前面的内容
		idx := strings.Index(line, ";")
		if idx > 0 {
			line = line[:idx]
		}

		tag := "base"
		lidx := strings.Index(line, "[")
		ridx := strings.Index(line, "]")
		if lidx > -1 && ridx > -1 {
			line = line[lidx+1:ridx]
			tag = line
			ini.IniMap[tag] = make(map[string]string)
		} else if lidx < 0 && ridx < 0  {
			cidx := strings.Index(line, "=")
			if cidx == -1 {
				return rStatus.StatusOk
			}
			ini.IniMap[tag][line[:cidx]] = line[cidx+1:]
		}
	}
	return rStatus.StatusOk
}

func (ini *IniParser)GetValue(tag, name string) (string, rStatus.RInfo) {
	m, ok := ini.IniMap[tag]
	if !ok {
		return "", rStatus.StatusError
	}
	v, ok := m[name]
	if !ok {
		return "", rStatus.StatusError
	}
	return v, rStatus.StatusOk
}
