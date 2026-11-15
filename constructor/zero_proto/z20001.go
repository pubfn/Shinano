package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_20001(msg *[]byte) ([]byte, int) {
	sc_19001 := &protobuf.SC_20001{
		Info: []*protobuf.TASKINFO{},
	}

	data, err := proto.Marshal(sc_19001)
	if err != nil {
		log.Fatal(err)
	}

	return data, 20001
}
