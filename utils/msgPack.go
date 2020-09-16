package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/whenfitrd/KServer/global"
	"github.com/whenfitrd/KServer/minterface"
)

func PackMsg(apiId, apiType int, data interface{}) []byte {
	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println("To JSON ERR:", err)
	}

	msgType, _ := IntToBytes(1,1)
	headInfo := []byte(global.MMsgHead)

	msgHead := append(msgType, headInfo...)

	dataLen, _ := IntToBytes(int(len(d)), 4)
	id, _ := IntToBytes(apiId, 4)
	t, _ := IntToBytes(apiType, 4)

	msg := append(msgHead, dataLen...)
	msg = append(msg, id...)
	msg = append(msg, t...)
	msg = append(msg, d...)

	return msg
}

func UnPackMsg(buffer []byte, msg minterface.IMsg) minterface.IMsg {
	msg.ParserHead(buffer[:global.MyMsgLen])
	buffer = buffer[global.MyMsgLen:]
	msg.ParserDataInfo(buffer[:global.MsgInfoLen])
	buffer = buffer[global.MsgInfoLen:]
	msg.Parser(buffer[:msg.GetMsgInfo().GetLength()])
	return msg
}

func IntToBytes(n int, b byte) ([]byte, error) {
	switch b {
	case 1:
		tmp := int8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 2:
		tmp := int16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 3, 4:
		tmp := int32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	}
	return nil, fmt.Errorf("IntToBytes b param is invaild")
}
