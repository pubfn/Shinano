package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_30001(msg *[]byte) ([]byte, int) {
	sc_30001 := &protobuf.SC_30001{
		UnreadNumber: proto.Uint32(0),
		TotalNumber:  proto.Uint32(2),
	}

	data, err := proto.Marshal(sc_30001)
	if err != nil {
		log.Fatal(err)
	}

	return data, 30001
}
