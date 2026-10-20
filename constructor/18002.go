package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_18002(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_18002{
		Score:               proto.Uint32(0),
		Rank:                proto.Uint32(1000),
		FightCount:          proto.Uint32(1),
		FightCountResetTime: proto.Uint32(1792436400),
		FlashTargetCount:    proto.Uint32(0),
		VanguardShipIdList:  []uint32{1},
		MainShipIdList:      []uint32{2},
		TargetList:          []*protobuf.TARGETINFO{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 18002
}
