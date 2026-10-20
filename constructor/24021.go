package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_24021(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_24021{
		Result:  proto.Uint32(0),
		Times:   []*protobuf.KVDATA{},
		Awards:  []*protobuf.KVDATA{},
		PassIds: []uint32{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 24021
}
