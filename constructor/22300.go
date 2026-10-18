package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_22300(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_22300{
		Handbooks: []*protobuf.TUTHANDBOOK{
			{Id: proto.Uint32(200601), Pt: proto.Uint32(1), Award: proto.Uint32(0), FinishedTaskIds: []uint32{}},
			{Id: proto.Uint32(100101), Pt: proto.Uint32(1), Award: proto.Uint32(0), FinishedTaskIds: []uint32{}},
			{Id: proto.Uint32(200501), Pt: proto.Uint32(1), Award: proto.Uint32(0), FinishedTaskIds: []uint32{}},
			{Id: proto.Uint32(200801), Pt: proto.Uint32(1), Award: proto.Uint32(0), FinishedTaskIds: []uint32{}},
			{Id: proto.Uint32(201401), Pt: proto.Uint32(1), Award: proto.Uint32(1), FinishedTaskIds: []uint32{}},
			{Id: proto.Uint32(200201), Pt: proto.Uint32(2), Award: proto.Uint32(0), FinishedTaskIds: []uint32{}},
		},
		FinishedTaskIds: []uint32{
			23004, 23003, 23048, 23039, 23026, 22011, 23035,
		},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 22300
}
