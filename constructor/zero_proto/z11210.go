package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11210(msg *[]byte) ([]byte, int) {
	sc_11200 := &protobuf.SC_11210{
		PermanentActivity: []uint32{},
		PermanentNow:      []uint32{},
	}

	data, err := proto.Marshal(sc_11200)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11210
}
