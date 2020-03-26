package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Test struct{
	Info string
	L int
}

type SliceMock struct {
	addr uintptr
	len int
	cap int
}

func main() {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	//inputReader := bufio.NewReader(os.Stdin)
	//fmt.Println("First")
	//clientName, _ := inputReader.ReadString('\n')
	//trimmedClient := strings.Trim(clientName, "\n")

	testStruct := &Test{"test struct.", 10000}
	testStruct1 := &Test{"test struct.", 5}

	d1, err := json.Marshal(testStruct)
	if err != nil {
		fmt.Println("To JSON ERR:", err)
	}

	d2, err := json.Marshal(testStruct1)
	if err != nil {
		fmt.Println("To JSON ERR:", err)
	}


	//Len := unsafe.Sizeof(*testStruct)
	//Len1 := unsafe.Sizeof(*testStruct)
	//testBytes := &SliceMock{
	//	addr: uintptr(unsafe.Pointer(testStruct)),
	//	cap: int(Len),
	//	len: int(Len),
	//}
	//data := *(*[]byte)(unsafe.Pointer(testBytes))
	//testBytes1 := &SliceMock{
	//	addr: uintptr(unsafe.Pointer(testStruct1)),
	//	cap: int(Len1),
	//	len: int(Len1),
	//}
	//data1 := *(*[]byte)(unsafe.Pointer(testBytes1))
	////buf := &bytes.Buffer{}
	////err = binary.Write(buf, binary.BigEndian, testStruct)
	////if err != nil {
	////	panic(err)
	////}
	////Len := len(buf.Bytes())
	//fmt.Println("head data length  ", len([]byte("use myMessage")))
	//
	////Len := unsafe.Sizeof(*testStruct)
	////data := *(*[]byte)(unsafe.Pointer(testStruct))
	////
	//////Len1 := unsafe.Sizeof(*testStruct1)
	////data1 := *(*[]byte)(unsafe.Pointer(testStruct1))
	//
	////len := unsafe.Sizeof("test: hello ...")
	//m2, _ := IntToBytes(int(len(data)), 4)
	//m3, _ := IntToBytes(1, 4)
	//
	//fmt.Println("data length  ", Len, len(data))
	//fmt.Println("data length  ", Len, len(data1))

	msgType, _ := IntToBytes(1,1)
	headInfo := []byte("use json")

	msgHead := append(msgType, headInfo...)

	m2, _ := IntToBytes(int(len(d1)), 4)
	m3, _ := IntToBytes(1, 4)

	m21, _ := IntToBytes(int(len(d2)), 4)
	m31, _ := IntToBytes(1, 4)

	//m4 := []byte("test: hello ...")
	msg := append(msgHead, m2...)
	msg = append(msg, m3...)
	msg = append(msg, d1...)

	msg1 := append(msgHead, m21...)
	msg1 = append(msg1, m31...)
	msg1 = append(msg1, d2...)
	_, err = conn.Write(msg)
	_, err = conn.Write(msg1)
	//conn.Close()
	//_, err = conn.Write([]byte("test: bye ..."))
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
