package msg

import (
	"github.com/golang/protobuf/proto"
	"github.com/whenfitrd/KServer/pb"
)

type PBMsg struct {
	MsgApiId int32
	Priority int32
	MsgData  []byte
}

func (pbm *PBMsg) GetPriority() int32 {
	return pbm.Priority
}

func (pbm *PBMsg) GetApiId() int32 {
	return pbm.MsgApiId
}

func (pbm *PBMsg) GetData() []byte {
	return pbm.MsgData
}

func (pbm *PBMsg) Parser(data []byte) {
	var baseInfo pb.Base

	proto.UnmarshalMerge(data, &baseInfo)

	pbm.MsgApiId = baseInfo.Head.Id
	pbm.Priority = baseInfo.Head.Priority
	pbm.MsgData = baseInfo.Info
}


