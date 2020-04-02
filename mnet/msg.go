package mnet

import (
	"encoding/binary"
	"github.com/whenfitrd/KServer/global"
	"github.com/whenfitrd/KServer/minterface"
	"github.com/whenfitrd/KServer/rStatus"
	"github.com/whenfitrd/KServer/utils"
	"strconv"
	"unsafe"
)

type Message struct {
	MsgType int8
	MsgInfo minterface.IMMsg
}

type MMsg struct {
	Length   int32
	MsgApiId int32
	MsgData  []byte
}

func (msg *Message) GetType() int8 {
	return msg.MsgType
}

func (msg *Message) GetMsgParser(msgType int8) minterface.IMMsg {
	switch msgType {
	case global.MyMessage:
		return &MMsg{}
	default:
		return &MMsg{}
	}
}

func (msg *Message) GetMsgInfo() minterface.IMMsg {
	return msg.MsgInfo
}

func (msg *Message) ParserHead(data []byte) rStatus.RStatus {
	logger.Info(string(data[1:]))
	defer utils.HandlePanic("server")
	if string(data[1:]) == global.MMsgHead {
		msg.MsgType = int8(data[0])
		logger.Info("End to read message head. type:" + strconv.Itoa(int(msg.MsgType)))
		return rStatus.Ok
	} else {
		logger.Error("Error about message head.")
		return rStatus.Error
	}
}

func (msg *Message) ParserDataInfo(data []byte) {
	mMsg := msg.GetMsgParser(msg.MsgType)

	mMsg.ParserDataInfo(data)

	msg.MsgInfo = mMsg
}

func (msg *Message) Parser(data []byte) {
	msg.MsgInfo.Parser(data)
}



func (msg *MMsg) GetLength() int32 {
	return msg.Length
}

func (msg *MMsg) GetApiId() int32 {
	return msg.MsgApiId
}

func (msg *MMsg) GetData() []byte {
	return msg.MsgData
}

func (msg *MMsg) ParserDataInfo(data []byte) {
	lenBuf := data[:unsafe.Sizeof(msg.Length)]
	msg.Length = int32(binary.BigEndian.Uint32(lenBuf))
	logger.Info("End to read message length. " + strconv.Itoa(int(msg.Length)))

	apiBuf := data[unsafe.Sizeof(msg.Length):unsafe.Sizeof(msg.Length)+unsafe.Sizeof(msg.MsgApiId)]
	msg.MsgApiId = int32(binary.BigEndian.Uint32(apiBuf))
	logger.Info("End to read message api id. " + strconv.Itoa(int(msg.MsgApiId)))
}

func (msg *MMsg) Parser(data []byte) {
	msg.MsgData = data
	logger.Info("End to read message data length. " + strconv.Itoa(int(len(msg.MsgData))))
}