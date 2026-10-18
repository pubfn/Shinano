package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_33114(msg *[]byte) ([]byte, int) {
	sc_33114 := &protobuf.SC_33114{
		IsWorldOpen: proto.Uint32(1),
		ShipIdList:  []uint32{},
		CmdIdList:   []uint32{},
		Progress:    proto.Uint32(0),
	}

	data, err := proto.Marshal(sc_33114)
	if err != nil {
		log.Fatal(err)
	}

	return data, 33114
}
