package zero_proto

import (
	"log"
	"time"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11000(msg *[]byte) ([]byte, int) {
	sc_10100 := &protobuf.SC_11000{
		Timestamp:               proto.Uint32(uint32(time.Now().Unix())),
		Monday_0OclockTimestamp: proto.Uint32(1606114800),
	}

	data, err := proto.Marshal(sc_10100)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11000
}
