package csText

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/whenfitrd/KServer/utils"
	"net"
)

type Test struct {
	Info string
	L    int
}

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func main() {
	//conn, err := net.Dial("tcp", "47.101.57.8:50000")
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	testStruct := &Test{"test struct.", 10000}

	m := utils.PackMsg(1, testStruct)

	_, err = conn.Write(m)
	conn.Close()
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
