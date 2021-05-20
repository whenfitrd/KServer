package msg

import (
	"encoding/binary"
	"github.com/whenfitrd/KServer/utils"
	"strconv"
	"unsafe"
)

type JsonMsg struct {
	MsgApiId int32
	Priority int32
	Length    int32
	MsgData  []byte
}

func (jm *JsonMsg) GetPriority() int32 {
	return jm.Priority
}

func (jm *JsonMsg) GetApiId() int32 {
	return jm.MsgApiId
}

func (jm *JsonMsg) GetLength() int32 {
	return jm.Length
}

func (jm *JsonMsg) GetData() []byte {
	return jm.MsgData
}

func (jm *JsonMsg) Parser(data []byte) {
	logger := utils.GetLogger()

	apiIdBuf := data[:unsafe.Sizeof(jm.MsgApiId)]
	jm.MsgApiId = int32(binary.BigEndian.Uint32(apiIdBuf))
	logger.Info("End to read message api id. " + strconv.Itoa(int(jm.MsgApiId)))

	priorityBuf := data[unsafe.Sizeof(jm.MsgApiId):unsafe.Sizeof(jm.MsgApiId)+unsafe.Sizeof(jm.Priority)]
	jm.Priority = int32(binary.BigEndian.Uint32(priorityBuf))
	logger.Info("End to read message priority. " + strconv.Itoa(int(jm.Priority)))

	lengthBuf := data[unsafe.Sizeof(jm.MsgApiId)+unsafe.Sizeof(jm.Priority):unsafe.Sizeof(jm.MsgApiId)+unsafe.Sizeof(jm.Priority)+unsafe.Sizeof(jm.Length)]
	jm.Length = int32(binary.BigEndian.Uint32(lengthBuf))
	logger.Info("End to read message length. " + strconv.Itoa(int(jm.Length)))

	jm.MsgData = data[unsafe.Sizeof(jm.MsgApiId)+unsafe.Sizeof(jm.Priority)+unsafe.Sizeof(jm.Length):]
}
