package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository/story"

	"google.golang.org/protobuf/proto"
)

func Build_40004(msg *[]byte) ([]byte, int) {
	cs := &protobuf.CS_40003{}
	if err := proto.Unmarshal(*msg, cs); err != nil {
		log.Fatal(err)
	}

	sc := &protobuf.SC_40004{
		Result:        proto.Uint32(0),
		DropInfo:      []*protobuf.DROPINFO{},
		ExtraDropInfo: []*protobuf.DROPINFO{},
		PlayerExp:     proto.Uint32(0),
		ShipExpList:   []*protobuf.SHIP_EXP{},
		Mvp:           proto.Uint32(story.GetMvp(cs.GetStatistics())),
		CommanderExp:  []*protobuf.COMMANDER_EXP{},
		HpDropInfo:    []*protobuf.HPDROPINFO{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 40004
}
