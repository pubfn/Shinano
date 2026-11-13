package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_12046(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_12046{
		Result: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12046
}
