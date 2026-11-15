package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_13201(msg *[]byte) ([]byte, int) {
	sc_13201 := &protobuf.SC_13201{
		CountList:             []*protobuf.EXPEDITION_DAILY_COUNT{},
		EliteExpeditionCount:  proto.Uint32(0),
		EscortExpeditionCount: proto.Uint32(0),
		ChapterCountList:      []*protobuf.EXPEDITION_DAILY_COUNT{},
		QuickExpeditionList:   []uint32{},
	}

	data, err := proto.Marshal(sc_13201)
	if err != nil {
		log.Fatal(err)
	}

	return data, 13201
}
