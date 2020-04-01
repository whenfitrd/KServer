package msg

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/whenfitrd/KServer/global"
)

func PackMsg(apiId int, data interface{}) []byte {
	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println("To JSON ERR:", err)
	}

	msgType, _ := IntToBytes(1,1)
	headInfo := []byte(global.MMsgHead)

	msgHead := append(msgType, headInfo...)

	dataLen, _ := IntToBytes(int(len(d)), 4)
	id, _ := IntToBytes(apiId, 4)

	msg := append(msgHead, dataLen...)
	msg = append(msg, id...)
	msg = append(msg, d...)

	return msg
}

func UnPackMsg(buffer []byte) *Message {
	msg := &Message{}
	msg.ParserHead(buffer[:global.MyMsgLen])
	buffer = buffer[global.MyMsgLen:]
	msg.MsgInfo = msg.ParserDataInfo(buffer[:global.MsgInfoLen]).(*MMsg)
	buffer = buffer[global.MsgInfoLen:]
	msg.Parser(buffer[:msg.MsgInfo.Length])
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
