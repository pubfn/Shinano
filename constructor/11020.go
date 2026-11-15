package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

// this looks like ping
func Build_11020(msg *[]byte) ([]byte, int) {
	sc_11018 := &protobuf.SC_11020{
		Result: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc_11018)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11020
}
