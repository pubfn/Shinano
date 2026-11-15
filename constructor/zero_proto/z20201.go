package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_20201(msg *[]byte) ([]byte, int) {
	sc_20201 := &protobuf.SC_20201{
		Info: []*protobuf.ACT_TASK_INIT_LIST{},
	}

	data, err := proto.Marshal(sc_20201)
	if err != nil {
		log.Fatal(err)
	}

	return data, 20201
}
