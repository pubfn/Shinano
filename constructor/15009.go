package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_15009(msg *[]byte) ([]byte, int) {
	sc_15001 := &protobuf.SC_15009{
		Result: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc_15001)
	if err != nil {
		log.Fatal(err)
	}

	return data, 15009
}
