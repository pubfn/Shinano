package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_27011(msg *[]byte) ([]byte, int) {
	sc_28000 := &protobuf.SC_27011{
		Endings:    []uint32{},
		Qualifieds: []uint32{},
	}

	data, err := proto.Marshal(sc_28000)
	if err != nil {
		log.Fatal(err)
	}

	return data, 27011
}
