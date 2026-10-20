package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_13506(msg *[]byte) ([]byte, int) {
	sc_13201 := &protobuf.SC_13506{
		RemapCountList: []*protobuf.REMAPCOUNT{},
	}

	data, err := proto.Marshal(sc_13201)
	if err != nil {
		log.Fatal(err)
	}

	return data, 13506
}
