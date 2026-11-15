package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_30101(msg *[]byte) ([]byte, int) {
	sc_30101 := &protobuf.SC_30101{
		Number:       proto.Uint32(0),
		MaxTimestamp: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc_30101)
	if err != nil {
		log.Fatal(err)
	}

	return data, 30101
}
