package global

import (
	"github.com/whenfitrd/KServer/rStatus"
	"log"

	"github.com/go-ini/ini"
)

func LoadIniFile(fileName string) (*ini.File, rStatus.RInfo) {
	cfgFile, err := ini.Load(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
		return nil, rStatus.StatusError
	}
	return cfgFile, rStatus.StatusOK
}
