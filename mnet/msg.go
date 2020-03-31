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
	MsgInfo *MMsg
}

type MMsg struct {
	Length   int32
	MsgApiId int32
	MsgData  []byte
}

func (msg *Message) GetType() int8 {
	return msg.MsgType
}

func (msg *Message) GetMsgParser(msgType int8) minterface.Msg {
	switch msgType {
	case global.MyMessage:
		return &MMsg{}
	default:
		return &MMsg{}
	}
}

func (msg *Message) ParserHead(data []byte) rStatus.RStatus {
	utils.GetLogger().Info(string(data[1:]))
	defer utils.HandlePanic("server")
	if string(data[1:]) == global.MMsgHead {
		msg.MsgType = int8(data[0])
		utils.GetLogger().Info("End to read message head. type:" + strconv.Itoa(int(msg.MsgType)))
		return rStatus.Ok
	} else {
		utils.GetLogger().Error("Error about message head.")
		return rStatus.Error
	}
}

func (msg *Message) ParserDataInfo(data []byte) minterface.Msg {
	mMsg := msg.GetMsgParser(msg.MsgType)

	mMsg.ParserDataInfo(data)

	return mMsg
}

func (msg *Message) Parser(data []byte) {
	////解析出信息的type
	//utils.GetLogger().Info("Start to read message type.")
	//typeBuf := make([]byte, unsafe.Sizeof(msg.MsgType))
	//if _, err := io.ReadFull(tc, typeBuf); err != nil {
	//	utils.GetLogger().Error("Error to read message type.")
	//	return
	//}
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
	utils.GetLogger().Info("End to read message length. " + strconv.Itoa(int(msg.Length)))

	apiBuf := data[unsafe.Sizeof(msg.Length):unsafe.Sizeof(msg.Length)+unsafe.Sizeof(msg.MsgApiId)]
	msg.MsgApiId = int32(binary.BigEndian.Uint32(apiBuf))
	utils.GetLogger().Info("End to read message api id. " + strconv.Itoa(int(msg.MsgApiId)))
}

func (msg *MMsg) Parser(data []byte) {
	//解析出信息的Length
	//lenBuf := make([]byte, unsafe.Sizeof(msg.Length))
	//if _, err := io.ReadFull(tc, lenBuf); err != nil {
	//	utils.GetLogger().Error("Error to read message length.", err)
	//	return
	//}
	//msg.Length = int32(binary.BigEndian.Uint32(lenBuf))
	//
	//utils.GetLogger().Info("End to read message length. " + strconv.Itoa(int(msg.Length)))
	//
	////解析出信息的MsgApiId
	//apiBuf := make([]byte, unsafe.Sizeof(msg.MsgApiId))
	//if _, err := io.ReadFull(tc, apiBuf); err != nil {
	//	utils.GetLogger().Error("Error to read message api id.", err)
	//	return
	//}
	//msg.MsgApiId = int32(binary.BigEndian.Uint32(apiBuf))
	//
	//utils.GetLogger().Info("End to read message api id. " + strconv.Itoa(int(msg.MsgApiId)))
	//
	////存入MsgData
	//dataBuf := make([]byte, msg.Length)
	////dataBuf := make([]byte, msg.Length)
	//if _, err := io.ReadFull(tc, dataBuf); err != nil {
	//	utils.GetLogger().Error("Error to read message data.", err)
	//	return
	//}
	//msg.MsgData = dataBuf
	//utils.GetLogger().Info("End to read message data length. " + strconv.Itoa(int(len(msg.MsgData))))

	msg.MsgData = data
	utils.GetLogger().Info("End to read message data length. " + strconv.Itoa(int(len(msg.MsgData))))
}