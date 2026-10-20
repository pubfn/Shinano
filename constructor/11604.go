package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11604(msg *[]byte) ([]byte, int) {
	sc_11604 := &protobuf.SC_11604{
		State:      proto.Uint32(0),
		SystemList: []uint32{},
		FailCount:  proto.Uint32(0),
		FailCd:     proto.Uint32(0),
		Notice:     proto.String(""),
	}

	data, err := proto.Marshal(sc_11604)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11604
}
