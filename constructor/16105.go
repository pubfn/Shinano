package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_16105(msg *[]byte) ([]byte, int) {
	sc_16105 := &protobuf.SC_16105{
		FirstPayList:    []uint32{},
		PayList:         []*protobuf.SHOPINFO{},
		NormalList:      []*protobuf.SHOPINFO{},
		NormalGroupList: []*protobuf.SHOPINFO{},
	}

	data, err := proto.Marshal(sc_16105)
	if err != nil {
		log.Fatal(err)
	}

	return data, 16105
}
