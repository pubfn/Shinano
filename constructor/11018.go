package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11018(msg *[]byte) ([]byte, int) {
	sc_11018 := &protobuf.SC_11018{
		Result:   proto.Uint32(0),
		DropList: []*protobuf.DROPINFO{},
	}

	data, err := proto.Marshal(sc_11018)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11018
}
