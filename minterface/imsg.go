package minterface

type Msg interface {
	GetLength() int32

	GetApiId() int32

	GetData() []byte

	Parser([]byte)

	ParserDataInfo([]byte)
}
