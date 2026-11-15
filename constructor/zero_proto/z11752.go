package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11752(msg *[]byte) ([]byte, int) {
	sc_11752 := &protobuf.SC_11752{
		Active: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc_11752)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11752
}
