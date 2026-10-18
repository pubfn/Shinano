package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_13002(msg *[]byte) ([]byte, int) {
	sc_13002 := &protobuf.SC_13002{
		CollectionList: []*protobuf.COLLECTIONINFO{
			{
				Id: proto.Uint32(30101), FinishTime: proto.Uint32(0),
				OverTime: proto.Uint32(0), ShipIdList: []uint32{},
			},
			{
				Id: proto.Uint32(30109), FinishTime: proto.Uint32(0),
				OverTime: proto.Uint32(0), ShipIdList: []uint32{},
			},
			{
				Id: proto.Uint32(30110), FinishTime: proto.Uint32(0),
				OverTime: proto.Uint32(0), ShipIdList: []uint32{},
			},
			{
				Id: proto.Uint32(30102), FinishTime: proto.Uint32(0),
				OverTime: proto.Uint32(0), ShipIdList: []uint32{},
			},
			{
				Id: proto.Uint32(20101), FinishTime: proto.Uint32(0),
				OverTime: proto.Uint32(0), ShipIdList: []uint32{},
			},
		},
		MaxTeam: proto.Uint32(2),
	}

	data, err := proto.Marshal(sc_13002)
	if err != nil {
		log.Fatal(err)
	}

	return data, 13002
}
