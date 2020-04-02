package global

import (
	"github.com/go-ini/ini"
)

var g *GlobalObj

func GetGObj() *GlobalObj {
	if g == nil {
		g = &GlobalObj{}
	}
	return g
}

type GlobalObj struct {
	IniFile *ini.File
}

func (g *GlobalObj)SetIniFile(iniFile *ini.File) {
	g.IniFile = iniFile
}
