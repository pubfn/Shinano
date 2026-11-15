package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_22300(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_22300{
		Handbooks:       []*protobuf.TUTHANDBOOK{},
		FinishedTaskIds: []uint32{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 22300
}
