package test

import (
	"github.com/whenfitrd/KServer/rStatus"
	"github.com/whenfitrd/KServer/utils"
	"testing"
)

func TestLoad(t *testing.T) {
	cParser := utils.GetConfigParser()
	rs := cParser.Load("./config.ini")
	if rs == rStatus.StatusOk {
		t.Log(cParser.IniMap)
	} else {
		t.Log("Load error")
	}
}

func TestGetValue(t *testing.T) {
	cParser := utils.GetConfigParser()
	cParser.Load("./config.ini")
	s, rs := cParser.GetValue("base", "ttt")
	if rs == rStatus.StatusOk {
		t.Log(s)
	} else {
		t.Log("GetValue error")
	}

}
