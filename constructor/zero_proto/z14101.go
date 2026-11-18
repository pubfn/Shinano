package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_14101(msg *[]byte) ([]byte, int) {
	sc_14101 := &protobuf.SC_14101{
		EquipSkinList: []*protobuf.EQUIPSKININFO{},
	}

	data, err := proto.Marshal(sc_14101)
	if err != nil {
		log.Fatal(err)
	}

	return data, 14101
}
