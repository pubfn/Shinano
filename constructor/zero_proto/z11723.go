package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11723(msg *[]byte) ([]byte, int) {
	sc_11711 := &protobuf.SC_11723{
		ResultList: []uint32{},
		OpTime:     proto.Uint32(0),
	}

	data, err := proto.Marshal(sc_11711)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11723
}
