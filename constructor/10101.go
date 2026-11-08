package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_10101(msg *[]byte) ([]byte, int) {
	sc_10100 := &protobuf.SC_10101{
		State: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc_10100)
	if err != nil {
		log.Fatal(err)
	}

	return data, 10101
}
